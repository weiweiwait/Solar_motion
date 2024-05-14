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

// 在 dao 方法中加入分页参数，并在查询中加入 Limit 和 Offset

func (dao *PrizeDao) GetActivePrizes(page int, pageSize int) (*[]model.Prize, error) {
	var prizes []model.Prize
	offset := (page - 1) * pageSize
	if err := dao.DB.Table("Prize").Where("end_date > ?", time.Now().Format("2006-01-02")).Offset(offset).Limit(pageSize).Find(&prizes).Error; err != nil {
		return nil, err
	}
	return &prizes, nil
}

//返回所有抽奖人的id

func (dao *PrizeDao) GetAllParticipantIds(prize_id uint) ([]uint, error) {
	var participantIds []uint
	err := dao.DB.Table("UserApply").Model(&model.Carry{}).Where("prize_id = ?", prize_id).Pluck("user_id", &participantIds).Error
	return participantIds, err
}
