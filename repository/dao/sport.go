package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
	"time"
)

type SportDao struct {
	*gorm.DB
}

func NewSportDao(ctx context.Context) *SportDao {
	return &SportDao{NewDBClient(ctx)}
}

func NewSportByDB(db *gorm.DB) *SportDao {
	return &SportDao{db}
}

//创建打卡记录

func (dao *SportDao) AddSportByUserId(sport *model.UserSport) (err error) {
	// 自动生成当前日期，这样每次调用这个函数时，我们都能获取当天的日期
	sport.SportDate = time.Now().UTC().Truncate(24 * time.Hour)

	// 在数据库中创建新打卡记录
	return dao.DB.Table("UserSport").Model(&model.UserSport{}).Create(&sport).Error
}
