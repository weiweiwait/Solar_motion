package model

type Carry struct {
	UserId uint `gorm:"column:user_id"`
	Name   string
}
