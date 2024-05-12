package user

import (
	"Solar_motion/config"
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/email"
	"Solar_motion/pkg/utils/jwt"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/pkg/utils/upload"
	"Solar_motion/repository/cache"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/model"
	"Solar_motion/types"
	"context"
	"errors"
	"mime/multipart"
	"sync"
	"time"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

//用户注册

func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserPhoneNumber(req.PhoneNumber)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New("该手机号已经注册")
		return
	}
	//加密密码
	user := &model.User{
		UserName:    req.UserName,
		QQ:          req.QQ,
		PhoneNumber: req.PhoneNumber,
		Integral:    0,
	}
	//加密密码
	if err = user.SetPassword(req.Password); err != nil {
		log.LogrusObj.Error(err)
		return
	}
	err = userDao.CreateUser(user)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}
func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserLoginReq) (resp interface{}, err error) {
	var user *model.User
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserPhoneNumber(req.PhoneNumber)
	if !exist {
		log.LogrusObj.Error(err)
		return nil, errors.New("用户不存在")
	}
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("账号/密码不正确")
	}
	accessToken, err := jwt.GenerateToken(user.ID, req.PhoneNumber)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	integral, _ := userDao.QueryIntegral(req.PhoneNumber)
	userResp := &types.UserLoginReply{
		PhoneNumber: req.PhoneNumber,
		Integral:    integral,
	}
	resp = &types.UserTokenData{
		User:        userResp,
		AccessToken: accessToken,
	}
	return
}

//修改头像

func (s *UserSrv) UserAvatarUpload(ctx context.Context, file multipart.File, fileSize int64) (resp interface{}, err error) {
	qConfig := config.Config.Oss
	ImgUrl := qConfig.QiNiuServer
	println(666)
	println(ImgUrl)
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	uId := u.Id
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	path, err := upload.ToQiNiu(file, fileSize)
	if err != nil {
		println(err)
	}
	user.Avatar = path
	err = userDao.UpdateUserAvatarById(uId, user.Avatar)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	resp = &types.UserAvatar{
		Avatar: path,
	}

	return

}

//发送邮件

func (s *UserSrv) UserSendEmail(ctx context.Context, req *types.UserSendEmail, accessToken string) (resp interface{}, err error) {

	key := "email:" + accessToken + ":" + req.QQ
	pan, _ := cache.RedisClient.Exists(cache.RedisContext, key).Result()
	if pan == 1 {
		expire, _ := cache.RedisClient.TTL(cache.RedisContext, key).Result()
		if expire > 120*time.Second {
			err := errors.New("请求频繁")
			return nil, err
		}
	}
	userDao := dao.NewUserDao(ctx)
	account, _ := userDao.UserExistsByqq(req.QQ)
	if account == false {
		err := errors.New("没有该用户")
		return nil, err
	}
	result := email.SendCode(req.QQ)
	if result == "" {
		err := errors.New("邮件发送失败，请检查邮件地址是否有效")
		return nil, err
	}
	err = cache.RedisClient.Set(cache.RedisContext, key, result, 3*time.Minute).Err()
	return nil, nil
}

//验证身份

func (s *UserSrv) ResetCodeVerify(ctx context.Context, req *types.UserSendCode, accessToken string) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	key := "email:" + accessToken + ":" + req.QQ
	exist, _ := cache.RedisClient.Exists(cache.RedisContext, key).Result()
	if exist == 0 {
		err := errors.New("请先请求一份邮件")
		return nil, err
	}
	value, err := cache.RedisClient.Get(cache.RedisContext, key).Result()
	if err != nil {
		return nil, err
	}
	if value == "" {
		err := errors.New("验证码失效，请重新请求")
		return nil, err
	}
	user1 := &model.User{
		Password: req.Password,
	}
	if err = user1.SetPassword(req.Password); err != nil {
		log.LogrusObj.Error(err)
		return
	}
	if value == req.Code {
		err = userDao.UpdateUserPasswordById(u.Id, user1.Password)
		if err != nil {
			return nil, err
		}
	} else {
		err := errors.New("验证码错误")
		return nil, err
	}
	return nil, nil
}

//修改名字

func (s *UserSrv) UpdateUserName(ctx context.Context, req *types.UserNameUpdate) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	err = userDao.UpdateUserNameById(u.Id, req.UserName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//查看所有获奖经理

func (s *UserSrv) GetAllPrice(ctx context.Context) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	history, err := userDao.GetAwardHistoryById(u.Id)
	resp = &types.UserHistory{
		AwardHistory: history,
	}
	return
}