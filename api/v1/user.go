package v1

import (
	"errors"
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.WithField("shouldBind", err).Error("shouldBind error")
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		// 校验参数
		if req.Key == "" || len(req.Key) != consts.EncryptMoneyKeyLength {
			log.LogrusObj.WithField("key", req.Key).Error("key长度错误,必须是6位数")
			err := errors.New("key长度错误,必须是6位数")
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		// 单例化操作
		l := service.GetUserSrv()
		resp, err := l.UserRegister(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
