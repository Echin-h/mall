package service

import (
	"context"
	"gin-mall/types"
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

func (s *PaymentSrv) Pay(ctx context.Context, req *types.PaymentDownReq) (resp interface{}, err error) {
	return
}
