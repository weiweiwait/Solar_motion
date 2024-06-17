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

// ExistOrNotByUserPhoneNumber 根据phone_name判断是否存在该名字
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

// 根据phone_number查询积分

func (dao *UserDao) QueryIntegral(phoneNumber string) (integral int, err error) {
	user := model.User{}
	err = dao.DB.Table("user").Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Integral, nil
}

//根据id返回user信息

func (dao *UserDao) GetUserById(uId uint) (user *model.User, err error) {
	err = dao.DB.Table("user").Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}

// UpdateUserById 根据 id 更新用户信息
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

// UpdateUserAvatarById 根据 id 更新用户信息

func (dao *UserDao) UpdateUserAvatarById(uId uint, avatar string) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(map[string]interface{}{"avatar": avatar}).Error
}

// 根据qq查询用户是否存在

func (dao *UserDao) UserExistsByqq(qq string) (exists bool, err error) {
	var count int64
	err = dao.DB.Table("user").Where("qq = ?", qq).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// UpdateUserPasswordById 根据 id 更新用户信息

func (dao *UserDao) UpdateUserPasswordById(uId uint, password string) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(map[string]interface{}{"password": password}).Error
}

// UpdateUserUerNameById 根据 id 更新用户信息

func (dao *UserDao) UpdateUserNameById(uId uint, username string) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(map[string]interface{}{"username": username}).Error
}

// GetAwardHistoryById 根据 id 获取用户的奖励历史
func (dao *UserDao) GetAwardHistoryById(uId uint) (awardHistory string, err error) {
	var user model.User
	err = dao.DB.Model(&model.User{}).Where("id=?", uId).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.AwardHistory, nil
}

// GetIntegralById 根据 id 获取用户的奖励历史
func (dao *UserDao) GetIntegralById(uId uint) (integral int, err error) {
	var user model.User
	err = dao.DB.Model(&model.User{}).Where("id=?", uId).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Integral, nil
}

// UpdateIntegralById 根据 id 更新用户信息

func (dao *UserDao) UpdateIntegralById(uId uint, integral int) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(map[string]interface{}{"integral": integral}).Error
}
func (dao *UserDao) GetUser(email string) *model.User1 {
	u := &model.User1{}
	dao.DB.Where("email = ?", email).Find(u)
	return u
}
