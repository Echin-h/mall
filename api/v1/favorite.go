package v1

import (
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListFavoritesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoritesServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		if req.PageSize == 0 {
			req.PageSize = consts.BasePageSize
		}

		l := service.GetFavoriteSrv()
		resp, err := l.FavoritesList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
