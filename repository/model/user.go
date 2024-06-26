package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	UserName     string `gorm:"column:username"`
	Password     string
	QQ           string
	PhoneNumber  string
	Avatar       string `gorm:"column:avatar"`
	Integral     int
	AwardHistory string `gorm:"type:text"`
}

const (
	PassWordCost        = 12       // 密码加密难度
	Active       string = "active" // 激活用户
)

// SetPassword 设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type User1 struct {
	gorm.Model
	NickName string `gorm:"type:varchar(25) not null"`
	Email    string `gorm:"type:varchar(20) not null unique"`
	Password string `gorm:"type:varchar(255)"`
	Icon     string `gorm:"default:default.jpg"`
}
