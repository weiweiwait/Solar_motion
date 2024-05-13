package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
)

type CarryDao struct {
	*gorm.DB
}

func NewCarryDao(ctx context.Context) *CarryDao {
	return &CarryDao{NewDBClient(ctx)}
}

func NewCarryDaoByDB(db *gorm.DB) *CarryDao {
	return &CarryDao{db}
}

// CreateCarry 创建抽奖记录
func (dao *CarryDao) CreateCarry(user *model.Carry) error {
	return dao.DB.Table("carry_prize").Model(&model.Carry{}).Create(&user).Error
}

//查看自己获奖记录

func (dao *CarryDao) QueryById(id uint) (user *[]model.Carry, err error) {
	err = dao.DB.Table("carry_prize").Model(&model.Carry{}).Where("user_id=?", id).
		Find(&user).Error
	return
}
