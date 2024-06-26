package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
)

type BlogDao struct {
	*gorm.DB
}

func NewBlogDao(ctx context.Context) *BlogDao {
	return &BlogDao{NewDBClient(ctx)}
}

func NewBlogByDB(db *gorm.DB) *BlogDao {
	return &BlogDao{db}
}

// CreateBlog 创建文章
func (dao *BlogDao) CreateBlog(blog *model.Blog) error {
	return dao.DB.Table("Blog").Model(&model.Blog{}).Create(&blog).Error
}

// CreatImages

func (dao *BlogDao) CreatImages(image *model.Images) error {
	return dao.DB.Table("Post_images").Model(&model.Images{}).Create(&image).Error
}

//返回文章所有图片

func (dao *BlogDao) GetImagesByUserIdAndBlogId(userid, blogId int) ([]model.Images, error) {
	var images []model.Images
	err := dao.DB.Table("Post_images").Where("user_id = ? AND blog_id = ?", userid, blogId).Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}

// GetAllBlogs 获取所有文章
func (dao *BlogDao) GetAllBlogs(page int, pageSize int) ([]model.Blog, error) {
	offset := (page - 1) * pageSize
	var blogs []model.Blog
	err := dao.DB.Table("Blog").Offset(offset).Limit(pageSize).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

//GetAllBlogImages 获取对应文章图片url
//func (dao *BlogDao) GetAllBlogImages(blog_id int , user_id int) ([]model.Images,error) {
//	var images []model.Images
//
//}
