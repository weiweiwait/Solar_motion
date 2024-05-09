package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Table("user").Model(&model.User{}).Create(&user).Error
}

// ExistOrNotByUserPhoneNumber 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserPhoneNumber(phoneNumber string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("phone_number = ?", phoneNumber).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}
