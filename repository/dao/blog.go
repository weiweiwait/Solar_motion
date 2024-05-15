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
