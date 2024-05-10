package api

import (
	"Solar_motion/consts"
	"Solar_motion/pkg/e"
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/service/user"
	"Solar_motion/types"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户注册

func RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		//参数检验
		if req.Password == "" {
			err := errors.New("密码不能为空")
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		l := user.GetUserSrv()
		resp, err := l.UserRegister(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

//用户根据手机号登录

func UserLoginHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.UserLoginReq
		if err := context.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusBadRequest, ErrorResponse(context, err))
			return
		}
		l := user.GetUserSrv()
		resp, err := l.UserLogin(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusInternalServerError, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

// 修改头像

func UserUpdateAvatar() gin.HandlerFunc {
	return func(context *gin.Context) {
		file, fileHeader, _ := context.Request.FormFile("file")
		if fileHeader == nil {
			err := errors.New(e.GetMsg(e.ErrorUploadFile))
			context.JSON(consts.IlleageRequest, ErrorResponse(context, err))
			log.LogrusObj.Infoln(err)
			return
		}
		fileSize := fileHeader.Size
		l := user.GetUserSrv()
		resp, err := l.UserAvatarUpload(context.Request.Context(), file, fileSize)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusInternalServerError, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

// 发送邮件
func SendEmail() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.UserSendEmail
		if err := context.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusBadRequest, ErrorResponse(context, err))
			return
		}
		accessToken := context.GetHeader("access_token")
		l := user.GetUserSrv()
		resp, err := l.UserSendEmail(context.Request.Context(), &req, accessToken)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusInternalServerError, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//用户通过邮箱验证身份并修改密码

func ResetCodeVerify() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.UserSendCode
		if err := context.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusBadRequest, ErrorResponse(context, err))
			return
		}
		accessToken := context.GetHeader("access_token")
		l := user.GetUserSrv()
		resp, err := l.ResetCodeVerify(context.Request.Context(), &req, accessToken)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusInternalServerError, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//修改用户名

func UserUpdateUserName() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.UserNameUpdate
		if err := context.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusBadRequest, ErrorResponse(context, err))
			return
		}
		l := user.GetUserSrv()
		resp, err := l.UpdateUserName(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusInternalServerError, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//查看所有获奖经理

func GetAllPrices() gin.HandlerFunc {
	return func(context *gin.Context) {
		l := user.GetUserSrv()
		resp, err := l.GetAllPrice(context.Request.Context())
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusInternalServerError, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}
