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
		//2.用户注册
		v1.POST("/user/register", api.RegisterHandler())
		//2.用户登录
		v1.POST("/user/login", api.UserLoginHandler())
	}
	return r
}
