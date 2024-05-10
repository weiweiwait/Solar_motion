package api

import (
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/service/integral"
	"github.com/gin-gonic/gin"
	"net/http"
)

//签到领取积分

func UserSignIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		l := integral.GetIntegralSrv()
		resp, err := l.SignIn(context.Request.Context())
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}
