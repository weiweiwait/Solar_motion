package manager

import (
	"Solar_motion/pkg/utils/jwt"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/repository/cache"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/model"
	"Solar_motion/types"
	"context"
	"errors"
	"sync"
	"time"
)

var ManagerSrvIns *ManagerSrv
var ManagerSrvOnce sync.Once

type ManagerSrv struct {
}

func GetManagerSrv() *ManagerSrv {
	ManagerSrvOnce.Do(func() {
		ManagerSrvIns = &ManagerSrv{}
	})
	return ManagerSrvIns
}

//管理员注册

func (s *ManagerSrv) ManagerRegister(ctx context.Context, req *types.ManagerRegisterReq) (resp interface{}, err error) {
	managerDao := dao.NewManagerDao(ctx)
	_, exist, err := managerDao.ExistOrNotByManagerPhoneNumber(req.PhoneNumber)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New("该手机号已经注册")
		return
	}
	manager := &model.Manager{
		UserName:    req.UserName,
		PhoneNumber: req.PhoneNumber,
	}
	//加密密码
	if err = manager.SetManagerPassword(req.Password); err != nil {
		log.LogrusObj.Error(err)
		return
	}
	err = managerDao.CreateManager(manager)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}

//管理员登录

func (s *ManagerSrv) ManagerLogin(ctx context.Context, req *types.ManagerLoginReq) (resp interface{}, err error) {
	var manager *model.Manager
	managerDao := dao.NewManagerDao(ctx)
	manager, exist, err := managerDao.ExistOrNotByManagerPhoneNumber(req.PhoneNumber)
	if !exist {
		log.LogrusObj.Error(err)
		return nil, errors.New("用户不存在")
	}
	if !manager.CheckManagerPassword(req.Password) {
		return nil, errors.New("账号/密码不正确")
	}
	accessToken, err := jwt.GenerateToken(manager.ID, req.PhoneNumber)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	managerResp := &types.ManagerLoginReply{
		PhoneNumber: req.PhoneNumber,
	}
	resp = &types.ManagerTokenData{
		User:        managerResp,
		AccessToken: accessToken,
	}
	return
}

// 管理员删除成员

func (s *ManagerSrv) ManagerDeleteUser(ctx context.Context, req *types.ManagerDeleteRep) (resp interface{}, err error) {
	managerDao := dao.NewManagerDao(ctx)
	err = managerDao.DeleteManagerUser(req.Username)
	if err != nil {
		return nil, err
	}
	return
}

//管理员发布抽奖活动

func (s *ManagerSrv) ManagerPushPrize(ctx context.Context, req *types.ManagerPushReq) (resp interface{}, err error) {
	managerDao := dao.NewPrizeDao(ctx)
	err = cache.RedisClient.Set(cache.RedisContext, req.Name, req.Sum, 24*time.Hour).Err()
	if err != nil {
		return nil, err
	}
	format := "2006-01-02"
	start, err := time.Parse(format, req.Start_Date)
	end, err := time.Parse(format, req.End_Date)
	prize := &model.Prize{
		Name:       req.Name,
		Describ:    req.Describ,
		Start_Date: start,
		End_Date:   end,
		Sum:        req.Sum,
		Status:     0,
	}
	err = managerDao.AddPrizeByManager(prize)
	if err != nil {
		return nil, err
	}
	resp = &types.ManagerPushReq{
		Name:       req.Name,
		Describ:    req.Describ,
		Start_Date: req.Start_Date,
		End_Date:   req.End_Date,
		Sum:        req.Sum,
	}
	return
}
