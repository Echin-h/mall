package v1

import (
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// CreateAddressHandler 新增收获地址
// @Summary 新增收获地址
// @Tags address
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param address body AddressCreateReq true "address"
// @Success 200 {object} ctl.Response{data=types.AddressCreateResp} "success"
// @Router /api/v1/addresses/create [post]
func CreateAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetAddressSrv()
		resp, err := l.AddressCreate(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// ShowAddressHandler 展示某个收获地址
// @Summary 展示某个收获地址
// @Tags address
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param id query string true "id"
// @Success 200 {object} ctl.Response{data=types.AddressGetResp} "success"
// @Router /api/v1/addresses/show [get]
func ShowAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressGetReq
		if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetAddressSrv()
		resp, err := l.AddressShow(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// ListAddressHandler 展示收获地址
// @Summary 展示收获地址
// @Tags address
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param page query string false "page"
// @Param pageSize query string false "pageSize"
// @Success 200 {object} ctl.Response{data=types.AddressListResp} "success"
// @Router /api/v1/addresses/list [get]
func ListAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressListReq
		if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetAddressSrv()
		resp, err := l.AddressList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// UpdateAddressHandler 修改收获地址
// @Summary 修改收获地址
// @Tags address
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param address body AddressServiceReq true "address"
// @Success 200 {object} ctl.Response{data=types.AddressServiceResp} "success"
// @Router /api/v1/addresses/update [post]
func UpdateAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressServiceReq
		if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetAddressSrv()
		resp, err := l.AddressUpdate(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// DeleteAddressHandler 删除收获地址
// @Summary 删除收获地址
// @Tags address
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param address body AddressDeleteReq true "address"
// @Success 200 {object} ctl.Response{data=types.AddressDeleteResp} "success"
// @Router /api/v1/addresses/delete [post]
func DeleteAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressDeleteReq
		if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetAddressSrv()
		resp, err := l.AddressDelete(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
