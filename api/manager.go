package api

import (
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/service/manager"
	"Solar_motion/types"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

//管理员注册

func ManagerRegister() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.ManagerRegisterReq
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		//参数检验
		if req.Password == "" {
			err := errors.New("密码不能为空")
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		l := manager.GetManagerSrv()
		resp, err := l.ManagerRegister(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//管理员登录

func ManagerLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.ManagerLoginReq
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		l := manager.GetManagerSrv()
		resp, err := l.ManagerLogin(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//管理员删除成员

func ManagerDeleteUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.ManagerDeleteRep
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		l := manager.GetManagerSrv()
		resp, err := l.ManagerDeleteUser(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//管理员发表奖品活动

func ManagerPushActivity() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.ManagerPushReq
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		l := manager.GetManagerSrv()
		resp, err := l.ManagerPushPrize(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}
