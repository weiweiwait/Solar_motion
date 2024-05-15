package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
)

type SeckillDao struct {
	*gorm.DB
}

func NewSeckillDao(ctx context.Context) *SeckillDao {
	return &SeckillDao{NewDBClient(ctx)}
}

func NewSeckillByDB(db *gorm.DB) *SeckillDao {
	return &SeckillDao{db}
}

// CreateCarryIntegral 创建抽积分记录

func (dao *SeckillDao) CreateCarryIntegral(user *model.Seckill) error {
	return dao.DB.Table("carry_integral").Model(&model.Seckill{}).Create(&user).Error
}

// 查询是否已经抢

func (dao *SeckillDao) UserExistsById(user_id, active_id int) (exists bool, err error) {
	var count int64
	err = dao.DB.Table("carry_integral").Where("user_id = ? AND active_id = ？", user_id, active_id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
