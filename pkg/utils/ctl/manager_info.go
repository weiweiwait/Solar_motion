package ctl

//
//import (
//	"context"
//	"errors"
//)
//
//type key1 int
//
//var managerKey key
//
//type ManagerInfo struct {
//	Id       uint   `json:"id"`
//	Username string `json:"username"`
//}
//
//func GetManagerInfo(ctx context.Context) (*ManagerInfo, error) {
//	manager, ok := FromContext(ctx)
//	if !ok {
//		return nil, errors.New("获取用户信息错误")
//	}
//	return manager, nil
//}
//
//func NewManagerContext(ctx context.Context, u *ManagerInfo) context.Context {
//	return context.WithValue(ctx, managerKey, u)
//}
//
//func FromManagerContext(ctx context.Context) (*ManagerInfo, bool) {
//	u, ok := ctx.Value(userKey).(*UserInfo)
//	return u, ok
//}
//
//func InitManagerInfo(ctx context.Context) {
//	// TOOD 放缓存，之后的用户信息，走缓存
//}
