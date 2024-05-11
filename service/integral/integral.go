package integral

import (
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/repository/cache"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/model"
	"context"
	"errors"
	"strconv"
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

//运动打卡

func (s *IntegralSrv) StartSport(ctx context.Context) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	println()
	time1 := time.Now().UTC().Truncate(24 * time.Hour)
	timeStr := time1.Format("2006-01-02")
	myUintAsString := strconv.Itoa(int(u.Id))
	key := myUintAsString + ":" + timeStr
	exist, _ := cache.RedisClient.Exists(cache.RedisContext, key).Result()
	if exist != 0 {
		err := errors.New("已经打卡")
		return nil, err
	}
	err = cache.RedisClient.Set(cache.RedisContext, key, u.Username, 24*time.Hour).Err()
	userDao := dao.NewUserDao(ctx)
	integral, err := userDao.GetIntegralById(u.Id)
	integral = integral + 10
	user := &model.User{
		Integral: integral,
	}
	err = userDao.UpdateIntegralById(u.Id, user.Integral)
	if err != nil {
		return nil, err
	}
	return
}
