package service

import (
	"context"
	"errors"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/respository/db/dao"
	"gin-mall/types"
	"sync"
)

var CartSrvIns *CartSrv
var CartSrvOnce sync.Once

type CartSrv struct{}

func GetCartSrv() *CartSrv {
	CartSrvOnce.Do(func() {
		CartSrvIns = &CartSrv{}
	})
	return CartSrvIns
}

func (s *CartSrv) CartCreate(ctx context.Context, req *types.CartCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	// 商品是否存在
	_, err = dao.NewProductDao(ctx).GetProductById(req.ProductId)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	cartDao := dao.NewCartDao(ctx)
	_, status, _ := cartDao.CreateCart(u.Id, req)
	if status == e.ErrorProductMoreCart {
		err = errors.New(e.GetMsg(status))
		return
	}

	return
}

func (s *CartSrv) ListCart(ctx context.Context, req *types.CartListReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	cartList, err := dao.NewCartDao(ctx).ListCart(u.Id, req)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	resp = &types.DataListResp{
		Item:  cartList,
		Total: int64(len(cartList)),
	}

	return
}
