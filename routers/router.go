package routers

import (
	"Solar_motion/api"
	"Solar_motion/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		//用户操作
		//1.用户注册
		v1.POST("/user/register", api.RegisterHandler())
		//2.用户登录
		v1.POST("/user/login", api.UserLoginHandler())
		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.AuthMiddleware())
		{
			//3.修改头像
			authed.POST("user/update", api.UserUpdateAvatar())
			//4.获取验证码
			authed.POST("/user/sendemail", api.SendEmail())
			//5.修改密码
			authed.POST("user/update_password", api.ResetCodeVerify())
			//6.修改用户名
			authed.POST("user/update_username", api.UserUpdateUserName())
		}
	}
	return r
}
