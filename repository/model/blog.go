package model

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	Id       uint
	UserId   uint `gorm:"column:user_id"`
	Title    string
	Contexts string `gorm:"type:text"`
}
type Images struct {
	BlogId uint   `gorm:"column:blog_id"`
	UserId uint   `gorm:"column:user_id"`
	Image  string `gorm:"column:images"`
}
type Blog1 struct {
	CreatedAt      time.Time `json:"createdAt"` //创建时间
	BlogId         string    `json:"blogId,omitempty"`
	Email          string    `json:"string"`
	Content        string    `json:"content"`
	BlogTitle      string    `json:"blogTitle"`
	Pictures       []string  `json:"pictures"`       // 图片
	GetLikesNumber int       `json:"getLikesNumber"` // 获赞数
	Location       string    `json:"location"`       //位置
}

type LikeBlog struct {
	gorm.Model
	UserID string `gorm:"column:user_id not null unique"` //用户id
	BlogId string `gorm:"colum:blog_id not nut unique"`   // 点赞文章Id
}

func (lBlog *LikeBlog) TableName() string {
	return "like_blog"
}
