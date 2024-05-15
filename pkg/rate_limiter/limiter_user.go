package rate_limiter

//
//import (
//	"context"
//	"fmt"
//	"golang.org/x/time/rate"
//	"sync"
//	"time"
//)
//
////用来在抢积分时对用户限流
//
//type LimiterHolder struct {
//	sync.RWMutex
//	visitors map[string]*rate.Limiter
//	max      int
//	interval time.Duration
//}
//
//func NewLimiterHolder(r rate.Limit, b int) *LimiterHolder {
//	return &LimiterHolder{
//		visitors: make(map[string]*rate.Limiter),
//		max:      b,
//		interval: time.Duration(float64(int(time.Second)) / float64(r)),
//	}
//}
//
//func (h *LimiterHolder) Take(ID string) (bool, error) {
//	h.Lock()
//	limiter, exists := h.visitors[ID]
//	if !exists {
//		limiter = rate.NewLimiter(rate.Every(h.interval), h.max)
//		h.visitors[ID] = limiter
//	}
//	h.Unlock()
//
//	// 尝试使用一个token，如果获取成功则返回 nil
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//	if limiter.Wait(ctx) == nil {
//		return true, nil
//	} else {
//		return false, fmt.Errorf("Too many requests from user %s", ID)
//	}
//}
//
//func cleanStaleEntries(limitHolder *LimiterHolder) {
//	for {
//		time.Sleep(time.Minute)
//		limitHolder.Lock()
//		for id, limiter := range limitHolder.visitors {
//			if limiter.Allow() == false {
//				delete(limitHolder.visitors, id)
//			}
//		}
//		limitHolder.Unlock()
//	}
//}
