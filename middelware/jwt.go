package middelware

import (
	"gin-mall/consts"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/jwt"
	"github.com/gin-gonic/gin"
)

// 鉴权的中间件 - 流程
// 1. 判断时候有token - 没有直接return
// 2. 解析token是否过期
//    - 如果没有过期，则会重新刷新token的时间
//    - 如果过期 |
// 3. 判断RefreshToken是否过期
//    - 如果没有过期，则重新生成token
//    - 如果过期，则返回错误
// 5. 通过token解析出用户信息
// 6. 将用户信息放入请求上下文
// 7. 继续执行

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = e.SUCCESS
		acceessToken := ctx.GetHeader("access_token")
		refreshToken := ctx.GetHeader("refresh_token")
		if acceessToken == "" {
			code = e.InvalidParams
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token不能为空",
			})
			ctx.Abort()
			return
		}
		newAccessToken, newRefreshToken, err := jwt.ParseRefreshToken(acceessToken, refreshToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		}
		if code != e.SUCCESS {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "鉴权失败",
				"error":  err.Error(),
			})
			ctx.Abort()
			return
		}
		claims, err := jwt.ParseToken(newAccessToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   err.Error(),
			})
			ctx.Abort()
			return
		}
		SetToken(ctx, newAccessToken, newRefreshToken)
		ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(), &ctl.UserInfo{Id: claims.ID}))
		ctl.InitUserInfo(ctx.Request.Context())
		ctx.Next()
	}
}

func SetToken(ctx *gin.Context, accessToken, refreshToken string) {
	//secure := IsHttps(ctx)
	// 把token既存于header又存于cookie
	ctx.Header("access_token", accessToken)
	ctx.Header("refresh_token", refreshToken)
	ctx.SetCookie("access_token", accessToken, consts.MaxAge, "/", "", false, true) //  secure 可以填写
	ctx.SetCookie("refresh_token", refreshToken, consts.MaxAge, "/", "", false, true)
}

func IsHttps(ctx *gin.Context) bool {
	if ctx.Request.TLS != nil || ctx.GetHeader(consts.HeaderForwardedProto) == "https" {
		return true
	}
	return false
}
