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

// CreateCartHandler 创建购物车
// @Summary 创建购物车
// @Tags cart
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param cart body CartCreateReq true "cart"
// @Success 200 {object} ctl.Response{data=types.CartCreateResp} "success"
// @Router /api/v1/cart/create [post]
func CreateCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CartCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetCartSrv()
		resp, err := l.CartCreate(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
		return
	}
}

// ListCartHandler 购物车详细信息
// @Summary 购物车详细信息
// @Tags cart
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {object} ctl.Response{data=types.CartListResp} "success"
// @Router /api/v1/cart/list [get]
func ListCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CartListReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		if req.PageSize == 0 {
			req.PageSize = consts.BasePageSize
		}

		l := service.GetCartSrv()
		resp, err := l.ListCart(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
		return
	}
}

// UpdateCartHandler 更新购物车
// @Summary 更新购物车
// @Tags cart
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param cart body UpdateCartServiceReq true "cart"
// @Success 200 {object} ctl.Response{data=types.UpdateCartServiceResp} "success"
// @Router /api/v1/cart/update [post]
func UpdateCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UpdateCartServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetCartSrv()
		resp, err := l.UpdateCart(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
		return

	}
}

// DeleteCartHandler 删除购物车
// @Summary 删除购物车
// @Tags cart
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param cart body DeleteCartReq true "cart"
// @Success 200 {object} ctl.Response{data=types.DeleteCartResp} "success"
// @Router /api/v1/cart/delete [post]
func DeleteCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteCartReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetCartSrv()
		resp, err := l.DeleteCart(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
