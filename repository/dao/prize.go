package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
	"time"
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

func (dao *PrizeDao) GetActivePrizes() (*[]model.Prize, error) {
	var prizes []model.Prize
	if err := dao.DB.Table("Prize").Where("end_date > ?", time.Now().Format("2006-01-02")).Find(&prizes).Error; err != nil {
		return nil, err
	}
	return &prizes, nil
}
