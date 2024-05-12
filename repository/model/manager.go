package model

import "golang.org/x/crypto/bcrypt"

type Manager struct {
	ID          uint
	UserName    string `gorm:"column:username"`
	Password    string
	PhoneNumber string
	Avatar      string `gorm:"column:avatar"`
}

const (
	ManagerPassWordCost = 12 // 密码加密难度
)

// SetPassword 设置密码

func (m *Manager) SetManagerPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), ManagerPassWordCost)
	if err != nil {
		return err
	}
	m.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码

func (m *Manager) CheckManagerPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}
