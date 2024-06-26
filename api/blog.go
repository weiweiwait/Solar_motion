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

//查看这篇文章的所有图片

func GetAllBlogImages() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.OtherImagesReq
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		l := blog.GetBlogSrv()
		resp, err := l.UserGetAllBlogImages(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

// 发布文章

func PostBlog() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.BlogService
		l := blog.GetBlogSrv()
		//file, fileHeader, _ := context.Request.FormFile("file")
		//if fileHeader == nil {
		//	err := errors.New(e.GetMsg(e.ErrorUploadFile))
		//	context.JSON(consts.IlleageRequest, ErrorResponse(context, err))
		//	log.LogrusObj.Infoln(err)
		//	return
		//}
		//fileSize := fileHeader.Size
		resp, err := l.PostBlog(context, &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

// 根据关键词搜索

func SearchBlog() gin.HandlerFunc {
	return func(context *gin.Context) {
		keyword := context.Param("keyword")
		page, err := strconv.Atoi(context.Param("page"))
		var req types.BlogService
		l := blog.GetBlogSrv()
		if keyword != "" && err == nil && page > 0 {
			resp, err := l.SearchByKeyWord(context, keyword, page, &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				context.JSON(http.StatusOK, ErrorResponse(context, err))
				return
			}
			context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
		}

	}
}

// 查询文章列表
func GetBlogList() gin.HandlerFunc {
	return func(context *gin.Context) {
		way := context.Param("way")
		page, err := strconv.Atoi(context.Param("page"))
		var req types.BlogService
		l := blog.GetBlogSrv()
		if way != "" && err == nil && page > 0 {
			resp, err := l.GetBlogList(context, way, page, &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				context.JSON(http.StatusOK, ErrorResponse(context, err))
				return
			}
			context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
		}
	}
}
