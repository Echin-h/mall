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

// ListFavoritesHandler 收藏列表
// @Summary 收藏列表
// @Tags favorite
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {object} ctl.Response{data=types.FavoritesServiceResp} "success"
// @Router /api/v1/favorites/list [get]
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

// CreateFavoritesHandler 新增收藏
// @Summary 新增收藏
// @Tags favorite
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param favorite body FavoriteCreateReq true "favorite"
// @Success 200 {object} ctl.Response{data=types.FavoriteCreateResp} "success"
// @Router /api/v1/favorites/create [post]
func CreateFavoritesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoriteCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		if req.ProductId == 0 || req.BossId == 0 {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, errors.New("params error")))
			return
		}

		l := service.GetFavoriteSrv()
		resp, err := l.FavoriteCreate(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// DeleteFavoritesHandler 删除收藏
// @Summary 删除收藏
// @Tags favorite
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param favorite body FavoriteDeleteReq true "favorite"
// @Success 200 {object} ctl.Response{data=types.FavoriteDeleteResp} "success"
// @Router /api/v1/favorites/delete [post]
func DeleteFavoritesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoriteDeleteReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetFavoriteSrv()
		resp, err := l.FavoriteDelete(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
