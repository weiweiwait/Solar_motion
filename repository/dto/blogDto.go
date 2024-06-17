package dto

import (
	"Solar_motion/repository/model"
	"time"
)

type UserDto struct {
	Email    string `json:"email,omitempty"`
	NickName string `json:"nickName,omitempty"`
	Icon     string `json:"icon"`
	Token    string `json:"token,omitempty"`
	//HasLikedBlog []string `json:"hasLikedBlog"`
}

func BuildUser(user *model.User1, token string) *UserDto {
	return &UserDto{
		Email:    user.Email,
		NickName: user.NickName,
		Token:    token,
		Icon:     user.Icon,
	}
}

func BuildUserList(user []model.User1) []*UserDto {
	var users []*UserDto
	for _, user := range user {
		userDto := &UserDto{
			Email:    user.Email,
			NickName: user.NickName,
			Icon:     user.Icon,
		}
		users = append(users, userDto)
	}
	return users
}

type BlogDto struct {
	CreatedAt      time.Time `json:"createdAt"` //创建时间
	BlogId         string    `json:"blogId,omitempty"`
	Email          string    `json:"email,omitempty"`
	Author         UserDto   `json:"author"` //作者
	Content        string    `json:"content"`
	BlogTitle      string    `json:"blogTitle"`
	Pictures       []string  `json:"pictures"`            // 图片
	GetLikesNumber int       `json:"getLikesNumber"`      // 获赞数
	Location       string    `json:"location"`            //位置
	Highlight      string    `json:"highlight,omitempty"` //高亮
}

func BuildBlogList(blogs []*model.Blog1, highLight []string, users []*UserDto) (blogDtos []*BlogDto) {
	for i, blog := range blogs {
		blogDto := &BlogDto{
			CreatedAt:      blog.CreatedAt,
			BlogId:         blog.BlogId,
			Author:         *users[i],
			Content:        blog.Content,
			BlogTitle:      blog.BlogTitle,
			Pictures:       blog.Pictures,
			GetLikesNumber: blog.GetLikesNumber,
			Location:       blog.Location,
			Highlight:      highLight[i],
		}
		blogDtos = append(blogDtos, blogDto)
	}
	return
}
