package model

import "time"

type Prize struct {
	Name       string
	Describ    string
	Start_Date time.Time `gorm:"column:start_date"`
	End_Date   time.Time `gorm:"column:end_date"`
	Sum        int
	Status     int
}
