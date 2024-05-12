package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
)

type ManagerDao struct {
	*gorm.DB
}

func NewManagerDao(ctx context.Context) *ManagerDao {
	return &ManagerDao{NewDBClient(ctx)}
}

func NewManagerDaoByDB(db *gorm.DB) *ManagerDao {
	return &ManagerDao{db}
}

// CreateManager 创建用户
func (dao *ManagerDao) CreateManager(manager *model.Manager) error {
	return dao.DB.Table("manager").Model(&model.Manager{}).Create(&manager).Error
}

// ExistOrNotByManagerPhoneNumber 根据phone_name判断是否存在该名字
func (dao *ManagerDao) ExistOrNotByManagerPhoneNumber(phoneNumber string) (manager *model.Manager, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Manager{}).Where("phone_number = ?", phoneNumber).Count(&count).Error
	if count == 0 {
		return manager, false, err
	}
	err = dao.DB.Model(&model.Manager{}).Where("phone_number = ?", phoneNumber).First(&manager).Error
	if err != nil {
		return manager, false, err
	}
	return manager, true, nil
}

//根据电话删除user

func (dao *ManagerDao) DeleteManagerUser(username string) error {
	return dao.DB.Table("user").Where("username = ?", username).Delete(model.Manager{}).Error
}
