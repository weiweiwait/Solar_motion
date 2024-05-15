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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"mime/multipart"
	"strconv"
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

//查询所有活动

func (s *UserSrv) GetAllPrize(ctx context.Context, page int, pageSize int) (resp interface{}, err error) {
	userDao := dao.NewPrizeDao(ctx)
	prizes, err := userDao.GetActivePrizes(page, pageSize)
	resp = prizes
	return
}

//报名活动

func (s *UserSrv) UserApply(ctx context.Context, req *types.UserApply) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	userDao := dao.NewApplyDao(ctx)
	userDao1 := dao.NewUserDao(ctx)
	Account, _ := userDao.ApplyExistsById(req.Id)
	integral, err := userDao1.QueryIntegral(u.Username)
	if integral < 20 {
		err = errors.New("积分不够，无法参加")
		return nil, err
	}
	if Account == true {
		err = errors.New("已经报名，不要重复报名")
		return nil, err
	}
	integral = integral - 20
	err = userDao1.UpdateIntegralById(u.Id, integral)
	if err != nil {
		return nil, err
	}
	apply := &model.UserApply{
		UserId:  u.Id,
		PrizeId: req.Id,
		Name:    req.Name,
	}
	err = userDao.CreateApply(apply)
	if err != nil {
		return nil, err
	}
	return
}

func (s *UserSrv) GetAllApplyActivity(ctx context.Context, page int, pageSize int) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	userDao := dao.NewApplyDao(ctx)
	prizes, err := userDao.GetActiveAllPrizes(page, pageSize, u.Id)
	resp = prizes
	return
}
func (s *UserSrv) GetAllPrizeAlready(ctx context.Context) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	userDao := dao.NewCarryDao(ctx)
	prizes, err := userDao.QueryById(u.Id)
	if err != nil {
		return nil, err
	}
	resp = prizes
	return
}

//用户抢积分

// 定义一个存储消息的结构体

type KafkaMessage struct {
	UserId   int    `json:"user_id"`   // 用户 ID
	ActiveId int    `json:"active_id"` // 活动 ID
	Name     string `json:"name"`
	Sum      int    `json:"sum"`
}

// 初始化Kafka的生产者
// kafkaWriter 创建一个新的 kafka writer

var kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
	Brokers: []string{"localhost:9092"}, //kafka集群地址
	Topic:   "solar",                    //topic
})

func (s *UserSrv) GrabPoints(ctx context.Context, req *types.Points) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	active_id := req.ActiveId
	key := strconv.Itoa(active_id) + req.Name + ":"
	integral := cache.RedisClient.Get(cache.RedisContext, key)
	integral1, err := integral.Result()
	integral2, err := strconv.Atoi(integral1)
	integral2 = integral2 + 100
	if integral2 <= 0 {
		err = errors.New("积分已经抢完")
		return nil, err
	}
	//exist, _ := cache.RedisClient.Exists(cache.RedisContext, key).Result()
	//if exist == 0 {
	//	err = errors.New("该活动已经结束")
	//	return nil, err
	//}
	// 生成实例并将其转为 json 数据
	kafkaMessage := KafkaMessage{
		UserId:   int(u.Id),
		ActiveId: req.ActiveId,
		Name:     req.Name,
		Sum:      100,
	}
	message, err := json.Marshal(kafkaMessage)
	if err != nil {
		return nil, err
	}
	err = kafkaWriter.WriteMessages(ctx, kafka.Message{Value: message})
	if err != nil {
		return nil, err
	}
	err = ProcessMessages(ctx, "solar")
	if err != nil {
		return nil, err
	}
	return
}

//消费

func ProcessMessages(ctx context.Context, topic string) (err error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       topic,
		Partition:   0,
		StartOffset: kafka.LastOffset,       // 从最新的开始
		MaxWait:     500 * time.Millisecond, // 最大等待时间
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		kafkaMsg := &KafkaMessage{}
		err = json.Unmarshal(m.Value, kafkaMsg)
		if err != nil {
			return err
		}

		err = HandleGrabPointsWithLock(ctx, kafkaMsg)
		if err != nil {
			return err
		}
	}
	r.Close()
	return
}

//处理抢积分逻辑

func HandleGrabPointsWithLock(ctx context.Context, msg *KafkaMessage) (err error) {
	//redis分布式锁
	lockSuccess, err := cache.RedisClient.SetNX(cache.RedisContext, strconv.Itoa(msg.UserId), msg.ActiveId, time.Second*3).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail", err)
		return errors.New("get lock fail")
	} else {
		fmt.Println("get lock success")
	}
	err = BattleIntegral(ctx, msg.UserId, msg.ActiveId, msg.Name)
	if err != nil {
		return err
	}
	value, _ := cache.RedisClient.Get(cache.RedisContext, strconv.Itoa(msg.UserId)).Result()
	println(value)
	_, err = cache.RedisClient.Del(cache.RedisContext, strconv.Itoa(msg.UserId)).Result()
	if err != nil {
		return err
	}
	return
}

func BattleIntegral(ctx context.Context, userId, activeId int, name string) (err error) {
	integralDao := dao.NewSeckillDao(ctx)
	exist, err := integralDao.UserExistsById(userId, activeId)
	if exist == true {
		err = errors.New("已经参与")
		return err
	}
	u, err := ctl.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	user := &model.Seckill{
		UserId:   userId,
		ActiveId: activeId,
		Name:     name,
	}
	err = integralDao.CreateCarryIntegral(user)
	if err != nil {
		return err
	}
	integral, err := userDao.QueryIntegral(u.Username)
	if err != nil {
		return err
	}
	err = userDao.UpdateIntegralById(u.Id, integral+10)
	if err != nil {
		return err
	}
	return
}
