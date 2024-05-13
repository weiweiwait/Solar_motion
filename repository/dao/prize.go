package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
)

type PrizeDao struct {
	*gorm.DB
}

func NewPrizeDao(ctx context.Context) *PrizeDao {
	return &PrizeDao{NewDBClient(ctx)}
}

func NewPrizeByDB(db *gorm.DB) *PrizeDao {
	return &PrizeDao{db}
}

//创建打卡记录

func (dao *PrizeDao) AddPrizeByManager(prize *model.Prize) (err error) {
	// 在数据库中创建新打卡记录
	return dao.DB.Table("Prize").Model(&model.Prize{}).Create(&prize).Error
}
