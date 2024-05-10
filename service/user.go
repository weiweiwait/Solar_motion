package service

import (
	"Solar_motion/config"
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/pkg/utils/jwt"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/pkg/utils/upload"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/model"
	"Solar_motion/types"
	"context"
	"errors"
	"mime/multipart"
	"sync"
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
