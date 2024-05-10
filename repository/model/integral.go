package model

import (
	"time"
)

type UserCheckin struct {
	UserID      uint      `gorm:"column:user_id"`
	CheckinDate time.Time `gorm:"column:checkin_date"`
}
