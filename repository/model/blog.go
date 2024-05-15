package model

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
