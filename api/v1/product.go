package v1

import (
	"errors"
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

// ListProductsHandler 展示商品列表
// @Summary 展示商品列表
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Param category_id query int false "category_id"
// @Param keyword query string false "keyword"
// @Param sort query string false "sort"
// @Success 200 {object} ctl.Response{data=types.ProductListResp} "success"
// @Router /api/v1/products/list [get]
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

// CreateProductHandler 新增商品
// @Summary 新增商品
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param product body ProductCreateReq true "product"
// @Success 200 {object} ctl.Response{data=types.ProductCreateResp} "success"
// @Router /api/v1/products/create [post]
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

// ShowProductHandler 展示商品
// @Summary 展示商品
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param id query string true "id"
// @Success 200 {object} ctl.Response{data=types.ProductShowResp} "success"
// @Router /api/v1/products/show [get]
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

// UpdateProjectHandler 更新商品
// @Summary 更新商品
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param product body ProductUpdateReq true "product"
// @Success 200 {object} ctl.Response{data=types.ProductUpdateResp} "success"
// @Router /api/v1/products/update [post]
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

// DeleteProductHandler 删除商品
// @Summary 删除商品
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param product body ProductDeleteReq true "product"
// @Success 200 {object} ctl.Response{data=types.ProductDeleteResp} "success"
// @Router /api/v1/products/delete [post]
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

// SearchProductsHandler 搜索商品
// @Summary 搜索商品
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Param keyword query string false "keyword"
// @Success 200 {object} ctl.Response{data=types.ProductSearchResp} "success"
// @Router /api/v1/products/search [get]
func SearchProductsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductSearchReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		if req.PageSize == 0 {
			req.PageSize = consts.BaseProductPageSize
		}

		l := service.GetProductSrv()
		resp, err := l.ProductSearch(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// ListProductImgHandler 展示商品图片
// @Summary 展示商品图片
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param id query string true "id"
// @Success 200 {object} ctl.Response{data=types.ListProductImgResp} "success"
// @Router /api/v1/products/img/list [get]
func ListProductImgHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListProductImgReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		if req.ID == 0 {
			err := errors.New("参数错误，不能为空")
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetProductSrv()
		resp, err := l.ProductImgList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// ListCategoryHandler 展示商品分类
// @Summary 展示商品分类
// @Tags category
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Success 200 {object} ctl.Response{data=types.CategoryListResp} "success"
// @Router /api/v1/category/list [get]
func ListCategoryHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		l := service.GetCategorySrv()
		resp, err := l.CategoryList(ctx.Request.Context())
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
