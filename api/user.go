package api

import (
	"Solar_motion/pkg/utils/log"
	"Solar_motion/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户注册

func RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK)
		}
	}
}
