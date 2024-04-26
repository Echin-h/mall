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

// CreateOrderHandler 创建订单
func CreateOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetOrderSrv()
		resp, err := l.OrderCreate(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
		return
	}
}

// ListOrdersHandler 订单列表
func ListOrdersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderListReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		if req.PageSize == 0 {
			req.PageSize = consts.BasePageSize
		}

		l := service.GetOrderSrv()
		resp, err := l.OrderList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
		return
	}
}

// ShowOrderHandler 订单详情
func ShowOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderShowReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetOrderSrv()
		resp, err := l.OrderShow(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
		return
	}
}

// DeleteOrderHandler 更新订单
func DeleteOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderDeleteReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetOrderSrv()
		resp, err := l.OrderDelete(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
		return
	}
}
