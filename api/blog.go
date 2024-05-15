package api

import (
	"Solar_motion/consts"
	"Solar_motion/pkg/e"
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/service/blog"
	"Solar_motion/types"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//发布文章

func UserPushBlog() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.Blog
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		l := blog.GetBlogSrv()
		resp, err := l.UserPushBlog(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//上传文章的照片返回url

func UserPushBlogPhoto() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.ImagesReq
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		file, fileHeader, _ := context.Request.FormFile("file")
		if fileHeader == nil {
			err := errors.New(e.GetMsg(e.ErrorUploadFile))
			context.JSON(consts.IlleageRequest, ErrorResponse(context, err))
			log.LogrusObj.Infoln(err)
			return
		}
		fileSize := fileHeader.Size
		l := blog.GetBlogSrv()
		resp, err := l.UserPushPhoto(context.Request.Context(), file, fileSize, &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

//查看文章

func GetAllBlog() gin.HandlerFunc {
	return func(context *gin.Context) {
		pageStr := context.Query("page")
		pageSizeStr := context.Query("pageSize")
		page, errPage := strconv.Atoi(pageStr)
		if errPage != nil || page < 1 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(pageSizeStr)
		l := blog.GetBlogSrv()
		resp, err := l.UserGetAllBlog(context.Request.Context(), page, pageSize)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}
