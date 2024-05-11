package middleware

import (
	"Solar_motion/consts"
	"Solar_motion/pkg/e"
	"Solar_motion/pkg/utils/ctl"
	util "Solar_motion/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware token验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		code = e.SUCCESS
		accessToken := context.GetHeader("access_token")
		if accessToken == "" {
			code = e.InvalidParams
			context.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token不能为空",
			})
			context.Abort()
			return
		}
		newAccessToken, err := util.ParseRefreshToken(accessToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		}
		if code != e.SUCCESS {
			context.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "鉴权失败",
				"error":  err.Error(),
			})
			context.Abort()
			return
		}
		claims, err := util.ParseToken(accessToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
			context.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   err.Error(),
			})
			context.Abort()
			return
		}
		SetToken(context, newAccessToken)
		context.Request = context.Request.WithContext(ctl.NewContext(context.Request.Context(), &ctl.UserInfo{Id: claims.ID, Username: claims.Username}))
		ctl.InitUserInfo(context.Request.Context())
		context.Next()
	}
}
func SetToken(c *gin.Context, accessToken string) {
	c.Header(consts.AccessTokenHeader, accessToken)
}
