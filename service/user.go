package service

import (
	"context"
	"errors"
	conf "gin-mall/conf/sql"
	"gin-mall/consts"
	"gin-mall/pkg/util/log"
	"gin-mall/respository/db/dao"
	"gin-mall/respository/db/model"
	"gin-mall/types"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct{}

// 单例化操作，保证了某个操作只能使用一次
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New("用户已经存在了")
		return
	}
	user := &model.User{
		NickName: req.NickName,
		UserName: req.UserName,
		Status:   model.Active,
		Money:    consts.UserInitMoney,
	}

	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		log.LogrusObj.Error(err)
		return
	}

	// 加密money
	money, err := user.EncryptMoney(req.Key)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	user.Money = money
	user.Avatar = consts.UserDefaultAvatarLocal
	if conf.Config.System.UploadModel == consts.UploadModelOss {
		user.Avatar = consts.UserDefaultAvatarOss
	}

	err = userDao.CreateUser(user)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return

}
