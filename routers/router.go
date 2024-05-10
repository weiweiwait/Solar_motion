package routers

import (
	"Solar_motion/api"
	"Solar_motion/middleware"
	"Solar_motion/pkg/rate_limiter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	// 限流中间件(滑动窗口限流)
	windowSize := int64(1000)   // 窗口大小为1000毫秒
	qps := int32(100)           // 最大QPS为100
	ctx := context.Background() // 初始背景上下文
	window := rate_limiter.NewSlidingWindow(windowSize, qps, ctx)
	r.Use(func(c *gin.Context) {
		if !window.TryQuery() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests, please try again later"})
			c.Abort()
			return
		}
		c.Next()
	})

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
			//7.查看所有获奖经理
			authed.GET("user/get_history", api.GetAllPrices())
		}
	}
	return r
}
