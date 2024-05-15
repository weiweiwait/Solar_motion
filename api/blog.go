package api

import (
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/service/blog"
	"Solar_motion/types"
	"github.com/gin-gonic/gin"
	"net/http"
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

	}
}
