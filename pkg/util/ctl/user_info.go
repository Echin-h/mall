package ctl

import (
	"context"
	"encoding/json"
	"errors"
	"gin-mall/respository/cache"
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

// InitUserInfo 用户信息走缓存
func InitUserInfo(ctx context.Context) {
	_ = CacheUserInfo(ctx.Value(userKey).(*UserInfo))
	return
}

func CacheUserInfo(u *UserInfo) error {
	cache.RedisClient.Set(cache.RedisContext, string(rune(userKey)), u, 0)
	return nil
}

func GetUserInfoFromCache(ctx context.Context) (*UserInfo, error) {
	res, err := cache.RedisClient.Get(cache.RedisContext, string(rune(userKey))).Result()
	b := []byte(res)
	var u *UserInfo
	err = json.Unmarshal(b, u)
	return u, err
}
