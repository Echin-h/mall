package v1

import (
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// OrderPaymentHandler 订单支付
// @Summary 订单支付
// @Tags payment
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param payment body PaymentDownReq true "payment"
// @Success 200 {object} ctl.Response{data=types.PaymentDownResp} "success"
// @Router /api/v1/payment/down [post]
func OrderPaymentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.PaymentDownReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetPaymentSrv()
		resp, err := l.Payment(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
