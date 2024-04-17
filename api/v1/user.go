package v1

import (
	"gin-mall/pkg/util/log"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserReqisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.WithField("shouldBind", err).Error("shouldBind error")
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
	}
}
