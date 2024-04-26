package service

import (
	"context"
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/respository/db/dao"
	"gin-mall/respository/db/model"
	"gin-mall/types"
	"sync"
)

var AddressSrvIns *AddressSrv
var AddressSrvOnce sync.Once

type AddressSrv struct{}

func GetAddressSrv() *AddressSrv {
	AddressSrvOnce.Do(func() {
		AddressSrvIns = &AddressSrv{}
	})
	return AddressSrvIns
}

func (s *AddressSrv) AddressCreate(ctx context.Context, req *types.AddressCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	// 地址可以重复，不判断是否重复
	addrDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  u.Id,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = addrDao.CreateAddress(address)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}

func (s *AddressSrv) AddressShow(ctx context.Context, req *types.AddressGetReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	addrDao := dao.NewAddressDao(ctx)
	address, err := addrDao.GetAddressByID(req.Id, u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return address, nil
}

// 所谓的假的分页实现，哈哈哈哈
func (s *AddressSrv) AddressList(ctx context.Context, req *types.AddressListReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	if req.PageSize == 0 {
		req.PageSize = consts.BasePageSize
	}

	addrDao := dao.NewAddressDao(ctx)
	addresses, err := addrDao.ListAddress(u.Id, req)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return addresses, nil
}

func (s *AddressSrv) AddressUpdate(ctx context.Context, req *types.AddressServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	addrDao := dao.NewAddressDao(ctx)
	name := req.Name
	phone := req.Phone
	address := req.Address
	id := req.Id
	err = addrDao.UpdateAddressByID(u.Id, id, name, phone, address)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}

func (s *AddressSrv) AddressDelete(ctx context.Context, req *types.AddressDeleteReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	addrDao := dao.NewAddressDao(ctx)
	err = addrDao.DeleteAddressByID(u.Id, req.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}
