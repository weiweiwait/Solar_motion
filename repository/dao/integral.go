package dao

import (
	"Solar_motion/repository/model"
	"context"
	"gorm.io/gorm"
	"time"
)

type IntegralDao struct {
	*gorm.DB
}

func NewIntegralDao(ctx context.Context) *IntegralDao {
	return &IntegralDao{NewDBClient(ctx)}
}

func NewIntegralByDB(db *gorm.DB) *IntegralDao {
	return &IntegralDao{db}
}

//创建签到记录

func (dao *IntegralDao) AddCheckinByUserId(checkin *model.UserCheckin) (err error) {
	// 自动生成当前日期，这样每次调用这个函数时，我们都能获取当天的日期
	checkin.CheckinDate = time.Now().UTC().Truncate(24 * time.Hour)

	// 在数据库中创建新的签到记录
	return dao.DB.Table("UserCheckin").Model(&model.UserCheckin{}).Create(&checkin).Error
}

// FindCheckinByUserId 查找指定用户的签到记录
func (dao *IntegralDao) FindCheckinByUserId(userId uint) (*model.UserCheckin, error) {
	// 创建签到记录实例用于存储查询结果
	var checkin model.UserCheckin

	// 执行查询并将结果存入实例
	err := dao.DB.Table("UserCheckin").Model(&model.UserCheckin{}).
		Where("user_id = ?", userId). // 添加查询条件
		Order("checkin_date desc").   // 根据签到日期逆序排序
		First(&checkin).              // 将查询结果存入实例
		Error

	return &checkin, err
}

// FindCheckinByDate 查找指定日期的签到记录
func (dao *IntegralDao) FindCheckinByDate(date time.Time) (int, error) {
	// 创建签到记录实例用于存储查询结果
	var checkin model.UserCheckin

	// 执行查询并将结果存入实例
	err := dao.DB.Table("UserCheckin").Model(&model.UserCheckin{}).
		Where("date(checkin_date) = date(?)", date). // 添加查询条件
		First(&checkin).                             // 将查询结果存入实例
		Error

	// 判断查询结果
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil // 没有找到记录，返回0
		}
		return 0, err // 查询过程出错，返回错误
	}

	return 1, nil // 成功找到记录，返回1
}
