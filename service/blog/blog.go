package blog

import (
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/pkg/utils/upload"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/model"
	"Solar_motion/types"
	"context"
	"mime/multipart"
	"strconv"
	"sync"
)

var BlogSrvIns *BlogSrv
var BlogSrvOnce sync.Once

type BlogSrv struct {
}

func GetBlogSrv() *BlogSrv {
	BlogSrvOnce.Do(func() {
		BlogSrvIns = &BlogSrv{}
	})
	return BlogSrvIns
}

//用户发博客

func (s *BlogSrv) UserPushBlog(ctx context.Context, req *types.Blog) (resp interface{}, err error) {
	userDao := dao.NewBlogDao(ctx)
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	blog := &model.Blog{
		UserId:   u.Id,
		Title:    req.Title,
		Contexts: req.Contexts,
	}
	err = userDao.CreateBlog(blog)
	if err != nil {
		return nil, err
	}
	return
}

//用户上传图片

func (s *BlogSrv) UserPushPhoto(ctx context.Context, file multipart.File, fileSize int64, req *types.ImagesReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	userDao := dao.NewBlogDao(ctx)
	path, err := upload.ToQiNiu(file, fileSize)
	if err != nil {
		println(err)
	}
	println(666)
	println(req.BlogId)
	blogId, err := strconv.Atoi(req.BlogId)
	user := &model.Images{
		BlogId: uint(blogId),
		UserId: u.Id,
		Image:  path,
	}
	err = userDao.CreatImages(user)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	resp = path

	return
}
func (s *BlogSrv) UserGetAllBlog(ctx context.Context, page, pageSize int) (resp interface{}, err error) {
	userDao := dao.NewBlogDao(ctx)
	blogs, err := userDao.GetAllBlogs(page, pageSize)
	resp = blogs
	return
}

//获取这篇文章的所有图片

func (s *BlogSrv) UserGetAllBlogImages(ctx context.Context, req *types.OtherImagesReq) (resp interface{}, err error) {
	userDao := dao.NewBlogDao(ctx)
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	blogId, err := strconv.Atoi(req.BlogId)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	images, err := userDao.GetImagesByUserIdAndBlogId(userId, blogId)
	resp = images
	return
}
