package blog

import (
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/pkg/utils/upload"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/dto"
	"Solar_motion/repository/model"
	"Solar_motion/types"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"sync"
	"time"
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
func (s *BlogSrv) PostBlog(ctx *gin.Context, req *types.BlogService) (resp interface{}, err error) {
	blogDao := dao.NewBlogDao1(ctx)
	//获取图片
	form, _ := ctx.MultipartForm()
	fileHeaders := form.File["pictures"]
	//校验图片类型和大小
	for _, header := range fileHeaders {
		//校验文件
		if header.Size > (8 << 18) {
		}
		if typ := header.Header.Get("Content-Type"); typ != "image/png" &&
			typ != "image/gif" &&
			typ != "image/jpeg" &&
			typ != "image/jpg" &&
			typ != "image/jfif" &&
			typ != "image/bmp" {
		}

	}
	var pictureUrls []string
	for _, header := range fileHeaders {
		file, err := header.Open()
		if err != nil {
			log.LogrusObj.Error(err)
			return nil, err
		}
		if url, err := upload.ToQiNiu(file, header.Size); err != nil {
			log.LogrusObj.Error(err)
			return nil, err
		} else {
			pictureUrls = append(pictureUrls, url)
		}
		_ = file.Close()
	}

	//封装model
	blog := &model.Blog1{
		CreatedAt:      time.Now(),
		Email:          req.Email,
		Location:       req.Location,
		BlogTitle:      req.BlogTitle,
		Content:        req.Content,
		Pictures:       pictureUrls,
		GetLikesNumber: 0,
	}
	//存储在es中
	_, err = blogDao.IndexBlog(blog)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	return
}

func (s *BlogSrv) SearchByKeyWord(ctx *gin.Context, keyword string, page int, req *types.BlogService) (resp interface{}, err error) {
	blogDao := dao.NewBlogDao1(ctx)
	userDao := dao.NewUserDao(ctx)
	err, blogs, _ := blogDao.SearchByKeyWord(keyword, page)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	if len(blogs) == 0 {
		err = errors.New("没有跟多的文章了")
		log.LogrusObj.Error(err)
		return nil, err
	}
	var emails []string
	var users []model.User1
	for _, blog := range blogs {
		emails = append(emails, blog.Email)
	}
	for _, email := range emails {
		users = append(users, *userDao.GetUser(email)) //userDao.GetUser(email)
	}
	//mysql查找作者信息
	userDtos := dto.BuildUserList(users)
	resp = dto.BuildBlogList(blogs, make([]string, len(blogs)+1), userDtos)
	return resp, nil
}

func (s *BlogSrv) GetBlogList(ctx *gin.Context, way string, page int, req *types.BlogService) (resp interface{}, err error) {
	blogDao := dao.NewBlogDao1(ctx)
	userDao := dao.NewUserDao(ctx)
	blogs, err := blogDao.GetBlogList(way, page)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	if len(blogs) == 0 {
		err = errors.New("没有跟多的文章了")
		log.LogrusObj.Error(err)
		return nil, err
	}
	var emails []string
	for _, blog := range blogs {
		emails = append(emails, blog.Email)
	}
	var users []model.User1
	for _, email := range emails {
		users = append(users, *userDao.GetUser(email)) //userDao.GetUser(email)
	}
	userDtos := dto.BuildUserList(users)
	resp = dto.BuildBlogList(blogs, make([]string, len(blogs)+1), userDtos)
	return resp, nil
}
