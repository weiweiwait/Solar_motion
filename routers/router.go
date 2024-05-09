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
		v1.POST("/user/register", api.RegisterHandler())
	}
	return r
}
