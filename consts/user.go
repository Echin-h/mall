package consts

const UserInitMoney = "10000"

const (
	EmailOperationBinding = iota + 1
	EmailOperationNoBinding
	EmailOperationUpdatePassword
)

var EmailOperationMap = map[uint]string{
	EmailOperationBinding:        "您正在绑定邮箱, 请点击链接确定身份 %s",
	EmailOperationNoBinding:      "您正在解邦邮箱, 请点击链接确定身份 %s",
	EmailOperationUpdatePassword: "您正在修改密码, 请点击链接校验身份 %s",
}

const (
	UserDefaultAvatarLocal = "avatar.JPG"
	UserDefaultAvatarOss   = "xxxxxxxxxxx"
)

const (
	UploadModelOss   = "oss"
	UploadModelLocal = "local"
)

const (
	AccessTokenHeader    = "access_token"
	RefreshTokenHeader   = "refresh_token"
	HeaderForwardedProto = "X-Forwarded-Proto"
	MaxAge               = 3600 * 24
)

const (
	SmtpEmailPort = 465
	EmailSubject  = "EchinMail"
)
