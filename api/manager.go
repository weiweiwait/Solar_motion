package api

import (
	"Solar_motion/consts"
	"Solar_motion/pkg/e"
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/service/manager"
	"Solar_motion/types"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

//管理员查看所有抽奖活动

func ManagerGetAllPrizes() gin.HandlerFunc {
	return func(context *gin.Context) {
		pageStr := context.Query("page")
		pageSizeStr := context.Query("pageSize")
		page, errPage := strconv.Atoi(pageStr)
		if errPage != nil || page < 1 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(pageSizeStr)
		l := manager.GetManagerSrv()
		resp, err := l.ManagerGetAllPrizes(context.Request.Context(), page, pageSize)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//管理员开奖

func ManagerSetPrizes() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.CarryPrize
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		l := manager.GetManagerSrv()
		resp, err := l.ManagerStartPrize(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//管理员上传头像

func ManagerUpdateAvatar() gin.HandlerFunc {
	return func(context *gin.Context) {
		file, fileHeader, _ := context.Request.FormFile("file")
		if fileHeader == nil {
			err := errors.New(e.GetMsg(e.ErrorUploadFile))
			context.JSON(consts.IlleageRequest, ErrorResponse(context, err))
			log.LogrusObj.Infoln(err)
			return
		}
		fileSize := fileHeader.Size
		l := manager.GetManagerSrv()
		resp, err := l.ManagerAvatarUpload(context.Request.Context(), file, fileSize)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusInternalServerError, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}
