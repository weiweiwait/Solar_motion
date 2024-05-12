package model

import "time"

type UserSport struct {
	UserID    uint      `gorm:"column:user_id"`
	SportDate time.Time `gorm:"column:sport_date"`
}
