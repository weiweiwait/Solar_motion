package api

import (
	"Solar_motion/types"
	"github.com/gin-gonic/gin"
)

//用户注册

func RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {

		}
	}
}
