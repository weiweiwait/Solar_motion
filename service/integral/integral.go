package integral

import (
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/model"
	"context"
	"errors"
	"sync"
	"time"
)

var IntegralSrvIns *IntegralSrv
var IntegralSrvOnce sync.Once

type IntegralSrv struct {
}

func GetIntegralSrv() *IntegralSrv {
	IntegralSrvOnce.Do(func() {
		IntegralSrvIns = &IntegralSrv{}
	})
	return IntegralSrvIns
}

var lock sync.Mutex // 定义互斥锁

func (s *IntegralSrv) SignIn(ctx context.Context) (resp interface{}, err error) {
	lock.Lock()         // 加锁
	defer lock.Unlock() // 解锁，使用defer可以确保在函数结束时解锁

	u, err := ctl.GetUserInfo(ctx)
	integralDao := dao.NewIntegralDao(ctx)
	userDao := dao.NewUserDao(ctx)
	integral, err := userDao.GetIntegralById(u.Id)
	integral = integral + 10
	user := &model.User{
		Integral: integral,
	}
	time := time.Now().UTC().Truncate(24 * time.Hour)
	status, err := integralDao.FindCheckinByDate(time)
	if status == 0 {
		err = userDao.UpdateIntegralById(u.Id, user.Integral)
		if err != nil {
			return nil, err
		}
		check := &model.UserCheckin{
			UserID: u.Id,
		}
		err = integralDao.AddCheckinByUserId(check)
		if err != nil {
			return nil, err
		}
	} else {
		err := errors.New("已经签到")
		return nil, err
	}
	return
}
