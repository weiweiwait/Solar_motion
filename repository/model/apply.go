package model

type UserApply struct {
	UserId  uint   `gorm:"column:user_id"`
	PrizeId uint   `gorm:"column:prize_id"`
	Name    string `gorm:"column:name"`
}
