package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
)

type ApplyDao struct {
	*gorm.DB
}

func NewApplyDao(ctx context.Context) *ApplyDao {
	return &ApplyDao{NewDBClient(ctx)}
}

func NewApplyByDB(db *gorm.DB) *ApplyDao {
	return &ApplyDao{db}
}

// CreateApply 创建报名抽奖记录

func (dao *ApplyDao) CreateApply(user *model.UserApply) error {
	return dao.DB.Table("UserApply").Model(&model.UserApply{}).Create(&user).Error
}

//根据user_id查看是否报名

func (dao *ApplyDao) ApplyExistsById(id uint) (exists bool, err error) {
	var count int64
	err = dao.DB.Table("UserApply").Where("prize_id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
