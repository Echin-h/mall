package router

import (
	api "gin-mall/api/v1"
	"gin-mall/middelware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middelware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		//用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		// 商品操作
		v1.GET("product/list", api.ListProductsHandler())
		v1.GET("product/show", api.ShowProductHandler())
		//v1.POST("product/search", api.SearchProductsHandler())
		//v1.GET("product/imgs/list", api.ListProductImgHandler()) // 商品图片
		//v1.GET("category/list", api.ListCategoryHandler())       // 商品分类
		//v1.GET("carousels", api.ListCarouselsHandler())          // 轮播图

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middelware.AuthMiddleWare())
		{
			// 用户操作
			authed.POST("user/update", api.UserUpdateHandler())
			authed.GET("user/show_info", api.ShowUserInfoHandler())
			authed.POST("user/send_email", api.SendEmailHandler())
			authed.GET("user/valid_email", api.ValidEmailHandler())
			//authed.POST("user/following", api.UserFollowingHandler())
			//authed.POST("user/unfollowing", api.UserUnFollowingHandler())
			//authed.POST("user/avatar", api.UploadAvatarHandler()) // 上传头像

			// 商品操作
			authed.POST("product/create", api.CreateProductHandler())
			authed.POST("product/update", api.UpdateProjectHandler())
			authed.POST("product/delete", api.DeleteProductHandler())

		}
	}
	return r
}
