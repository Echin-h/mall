package v1

import (
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitSkillProductHandler 初始化秒杀商品信息
// @Summary 初始化秒杀商品信息
// @Tags skill_product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param skill_product body SkillProductReq true "skill_product"
// @Success 200 {object} ctl.Response{data=types.SkillProductResp} "success"
// @Router /api/v1/skill_product/init [post]
func InitSkillProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SkillProductReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetSkillProductSrv()
		resp, err := l.InitSkillProduct(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// ListSkillProductHandler 初始化秒杀商品信息
// @Summary 初始化秒杀商品信息
// @Tags skill_product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Success 200 {object} ctl.Response{data=types.ListSkillProductResp} "success"
// @Router /api/v1/skill_product/list [get]
func ListSkillProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListSkillProductReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetSkillProductSrv()
		resp, err := l.ListSkillGoods(ctx.Request.Context())
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// GetSkillProductHandler 获取秒杀商品的详情
// @Summary 获取秒杀商品的详情
// @Tags skill_product
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param id query string true "id"
// @Success 200 {object} ctl.Response{data=types.GetSkillProductResp} "success"
// @Router /api/v1/skill_product/get [get]
func GetSkillProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.GetSkillProductReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetSkillProductSrv()
		resp, err := l.GetSkillGoods(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// SkillProductHandler 秒杀商品
// @Summary 秒杀商品
// @Tags skill_product
// @Accept json
// @Produce json
// @Param Authorization header string true "
// @Param skill_product body SkillProductReq true "skill_product"
// @Success 200 {object} ctl.Response{data=types.SkillProductResp} "success"
// @Router /api/v1/skill_product [post]
func SkillProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SkillProductReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetSkillProductSrv()
		resp, err := l.SkillProduct(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
