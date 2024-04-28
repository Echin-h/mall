package service

import (
	"context"
	"errors"
	"fmt"
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/respository/db/dao"
	"gin-mall/respository/db/model"
	"gin-mall/types"
	"gorm.io/gorm"
	"sync"
)

var PaymentSrvIns *PaymentSrv
var PaymentSrvOnce sync.Once

type PaymentSrv struct{}

func GetPaymentSrv() *PaymentSrv {
	PaymentSrvOnce.Do(func() {
		PaymentSrvIns = &PaymentSrv{}
	})
	return PaymentSrvIns
}

func (s *PaymentSrv) Payment(ctx context.Context, req *types.PaymentDownReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	user, err := dao.NewUserDao(ctx).GetUserById(u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	rel := user.CheckPassword(req.Password)
	if !rel {
		err = errors.New("password error")
		log.LogrusObj.Error(err)
		return nil, err
	}

	err = dao.NewOrderDao(ctx).Transaction(func(tx *gorm.DB) error {
		uid := u.Id

		payment, err := dao.NewOrderDaoByDB(tx).GetOrderById(req.OrderId, uid)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

		money := payment.Money
		num := payment.Num
		money = money * float64(num)

		userDao := dao.NewUserDaoByDB(tx)
		user, err := userDao.GetUserById(uid)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

		// Decrypt the money. Subtract the order. Then encrypt it again.
		moneyFloat, err := user.DecryptMoney(req.Key)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}
		if moneyFloat-money < 0.0 {
			err = errors.New("insufficient money")
			log.LogrusObj.Error(err)
			return err
		}

		finMoney := fmt.Sprintf("%f", moneyFloat-money)
		user.Money, err = user.EncryptMoney(finMoney)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

		err = userDao.UpdateUserById(uid, user)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

		productDao := dao.NewProductDao(ctx)
		product, err := productDao.GetProductById(payment.ProductID)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}
		product.Num = product.Num - payment.Num
		err = productDao.UpdateProductById(payment.ProductID, product)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

		payment.Type = consts.OrderTypePendingShipping
		err = dao.NewOrderDaoByDB(tx).UpdateOrderById(req.OrderId, payment)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

		productUser := &model.Product{
			Name:          product.Name,
			CategoryID:    product.CategoryID,
			Title:         product.Title,
			Info:          product.Info,
			ImgPath:       product.ImgPath,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			Num:           num,
			OnSale:        false,
			BossID:        uid,
			BossName:      user.UserName,
			BossAvatar:    user.Avatar,
		}

		err = productDao.CreateProduct(productUser)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

		return nil
	})

	log.LogrusObj.Info("Payment success")
	return nil, err
}
