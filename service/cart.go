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

// 其实应该检测一下 num 是否大于 numMax

func (s *CartSrv) UpdateCart(ctx context.Context, req *types.UpdateCartServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	err = dao.NewCartDao(ctx).UpdateCartById(u.Id, req)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}

func (s *CartSrv) DeleteCart(ctx context.Context, req *types.DeleteCartReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	err = dao.NewCartDao(ctx).DeleteCartById(u.Id, req.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}
