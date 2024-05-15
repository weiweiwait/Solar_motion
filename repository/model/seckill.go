package model

type Seckill struct {
	UserId   int `gorm:"column:user_id"`
	ActiveId int `gorm:"column:active_id"`
	Name     string
}
