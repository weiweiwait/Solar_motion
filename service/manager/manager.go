package manager

import (
	"Solar_motion/pkg/utils/log"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/model"
	"Solar_motion/types"
	"context"
	"errors"
	"sync"
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
