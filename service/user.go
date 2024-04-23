package service

import (
	"context"
	"errors"
	"fmt"
	conf "gin-mall/conf/sql"
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/email"
	"gin-mall/pkg/util/jwt"
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

func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	var user *model.User
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if !exist {
		log.LogrusObj.Error("用户不存在", err)
		err = errors.New("用户不存在")
		return
	}

	if !user.CheckPassword(req.Password) {
		log.LogrusObj.Error("密码错误", err)
		err = errors.New("密码错误")
		return
	}
	token, refreshToken, err := jwt.GenerateToken(user.ID, req.UserName)
	if err != nil {
		log.LogrusObj.Error("生成token失败", err)
		return nil, err
	}

	userResp := &types.UserInfoResp{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	resp = &types.UserLoginResp{
		User:         userResp,
		Token:        token,
		RefreshToken: refreshToken,
	}

	return

}

func (s *UserSrv) UserUpdate(ctx context.Context, req *types.UserInfoUpdateReq) (resp interface{}, err error) {
	u, _ := ctl.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	if req.NickName != "" {
		user.NickName = req.NickName
	}

	err = userDao.UpdateUserById(u.Id, user)
	log.LogrusObj.Infoln(err)

	return
}

func (s *UserSrv) UserInfoShow(ctx context.Context, req *types.UserInfoShowReq) (resp interface{}, err error) {
	u, _ := ctl.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	resp = &types.UserInfoResp{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	return
}

func (s *UserSrv) SendEmail(ctx context.Context, req *types.SendEmailServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	var address string
	token, err := jwt.GenerateEmailToken(u.Id, req.Email, req.Password, req.OperationType)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	sender := email.NewEmailSender()
	address = conf.Config.Email.ValidEmail + token
	mailText := fmt.Sprintf(consts.EmailOperationMap[req.OperationType], address)
	if err = sender.Send(mailText, req.Email, consts.EmailSubject); err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}

func (s *UserSrv) ValidEmail(ctx context.Context, req *types.ValidEmailReq) (resp interface{}, err error) {
	var uId uint
	var eml string
	var pass string
	var OperationType uint

	if req.Token == "" {
		err = errors.New("token不存在")
		log.LogrusObj.Error(err, "skdksdkkdfs")
		return
	}

	claims, err := jwt.ParseEmailToken(req.Token)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	} else {
		uId = claims.ID
		eml = claims.Email
		pass = claims.Password
		OperationType = claims.OperationType
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	switch OperationType {
	case consts.EmailOperationBinding:
		user.Email = eml
	case consts.EmailOperationNoBinding:
		user.Email = ""
	case consts.EmailOperationUpdatePassword:
		err = user.SetPassword(pass)
		if err != nil {
			log.LogrusObj.Error("密码加密错误")
			return
		}
	default:
		return nil, errors.New("操作类型错误")
	}

	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	resp = &types.UserInfoResp{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	return

}
