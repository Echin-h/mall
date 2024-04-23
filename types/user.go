package types

type UserRegisterReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

type UserServiceReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

type UserInfoResp struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

type UserLoginResp struct {
	User         interface{} `json:"user"`
	Token        string      `json:"token"`
	RefreshToken string      `json:"refresh_token"`
}

type UserInfoUpdateReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	// 很好奇为什么只能改昵称。。。
	// 很好奇是不是
	// 我也很好奇啊
}

type UserInfoShowReq struct{}

type SendEmailServiceReq struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	// OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ValidEmailReq struct {
	Token string `json:"token" form:"token"`
}

type UserFollowingReq struct {
	Id uint `json:"id" form:"id"`
}
