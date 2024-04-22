package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

// 简单来说 GetUserInfo+FromContext 就是这个函数
//func SimplierGet (ctx context.Context)(*UserInfo,bool){
//	value, ok := ctx.Value(userKey).(*UserInfo)
//	if !ok {
//		return nil, false
//	}
//	return value, true
//}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return user, nil
}

func InitUserInfo(ctx context.Context) {
	// TOOD 放缓存，之后的用户信息，走缓存
	// 往缓存中存放 UserInfo这个结构体
}
