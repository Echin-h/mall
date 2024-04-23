package v1

import (
	"fmt"
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func ListProductsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductListReq
		if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		if req.PageSize == 0 {
			req.PageSize = consts.BaseProductPageSize
		}

		l := service.GetProductSrv()
		resp, err := l.ProductList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// TODO: create from both form and json is fail

func CreateProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductCreateReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		// 传入多个图片,一般来说 form 与 json 不要写在一起
		form, err := ctx.MultipartForm()
		if err != nil {
			fmt.Println("MultipartForm error : ", err)
		}
		files := form.File["image"]

		l := service.GetProductSrv()
		resp, err := l.ProductCreate(ctx.Request.Context(), files, &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

func ShowProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductShowReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetProductSrv()
		resp, err := l.ProductShow(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// TODO: 先实现用户信息缓存到Redis，然后在service层调用Redis，验证用户信息

func UpdateProjectHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductUpdateReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.LogrusObj.Error("json parse error")
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		l := service.GetProductSrv()
		resp, err := l.UpdateProduct(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error("UpdateProduct is wrong")
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

func DeleteProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductDeleteReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		l := service.GetProductSrv()
		resp, err := l.ProductDelete(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
