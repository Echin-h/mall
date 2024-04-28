package v1

import (
	"errors"
	"gin-mall/consts"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// UserRegisterHandler 用户注册
// @Summary 用户注册
// @Tags user
// @Accept json
// @Produce json
// @Param user body UserRegisterReq true "user"
// @Success 200 {object} ctl.Response
// @Router /api/v1/user/register [post]
func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.WithField("shouldBind", err).Error("shouldBind error")
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		// 校验参数
		if req.Key == "" || len(req.Key) != consts.EncryptMoneyKeyLength {
			log.LogrusObj.WithField("key", req.Key).Error("key长度错误,必须是6位数")
			err := errors.New("key长度错误,必须是6位数")
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		// 单例化操作
		l := service.GetUserSrv()
		resp, err := l.UserRegister(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// UserLoginHandler 用户登录
// @Summary 用户登录
// @Tags user
// @Accept json
// @Produce json
// @Param user body UserServiceReq true "user"
// @Success 200 {object} ctl.Response
// @Router /api/v1/user/login [post]
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.WithField("shouldBind", err).Error("shouldBind error")
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		// 单例化操作
		l := service.GetUserSrv()
		resp, err := l.UserLogin(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// UserUpdateHandler 用户信息更新
// @Summary 用户信息更新
// @Tags user
// @Accept json
// @Produce json
// @Param user body UserInfoUpdateReq true "user"
// @Success 200 {object} ctl.Response
// @Router /api/v1/user/update [post]
func UserUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoUpdateReq
		if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		// 单例化操作
		l := service.GetUserSrv()
		resp, err := l.UserUpdate(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// ShowUserInfoHandler 用户信息展示
// @Summary 用户信息展示
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param id query string true "id"
// @Success 200 {object} ctl.Response
// @Router /api/v1/user/show [get]
func ShowUserInfoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoShowReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		l := service.GetUserSrv()
		resp, err := l.UserInfoShow(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// SendEmailHandler 发送邮件
// @Summary 发送邮件
// @Tags user
// @Accept json
// @Produce json
// @Param email body SendEmailServiceReq true "email"
// @Success 200 {object} ctl.Response
// @Router /api/v1/user/send_email [post]
func SendEmailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SendEmailServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		l := service.GetUserSrv()
		resp, err := l.SendEmail(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))

	}
}

// ValidEmailHandler 邮箱验证
// @Summary 邮箱验证
// @Tags user
// @Accept json
// @Produce json
// @Param email body ValidEmailReq true "email"
// @Success 200 {object} ctl.Response
// @Router /api/v1/user/valid_email [post]
func ValidEmailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ValidEmailReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		t := ctx.GetHeader("token")
		req.Token = t
		l := service.GetUserSrv()
		resp, err := l.ValidEmail(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// UserFollowingHandler 用户关注
// @Summary 用户关注
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param user_id body UserFollowingReq true "user_id"
// @Success 200 {object} ctl.Response
// @Router /api/v1/user/following [post]
func UserFollowingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserFollowingReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		resp, err := l.UserFollow(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// UserUnFollowingHandler 不关注用户
// @Summary 不关注用户
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param user_id body UserFollowingReq true "user_id"
// @Success 200 {object} ctl.Response
// @Router /api/v1/user/un_following [post]
func UserUnFollowingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserFollowingReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		resp, err := l.UserUnFollow(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}

		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// UploadAvatarHandler 上传头像
// @Summary 上传头像
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "token<access_token><refresh_token>"
// @Param file formData file true "file"
// @Success 200 {object} ctl.Response
// @Router /api/v1/user/avatar [post]
func UploadAvatarHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, fileHeader, _ := ctx.Request.FormFile("file")
		if fileHeader == nil {
			err := errors.New(e.GetMsg(e.ErrorUploadFile))
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		resp, err := l.UserAvatarUpload(ctx.Request.Context(), file)
		if err != nil {
			log.LogrusObj.Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
