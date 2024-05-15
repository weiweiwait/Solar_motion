package kafka

//import (
//	"Solar_motion/repository/cache"
//	"context"
//	"encoding/json"
//	"github.com/segmentio/kafka-go"
//)
//
//// 声明一个用于反序列化消息的结构体定义
//
//type Request struct {
//	UserID   string `json:"user_id"`
//	ActiveID string `json:"active_id"`
//}
//
//func ConsumeMessage(req *Request) {
//	r := kafka.NewReader(kafka.ReaderConfig{
//		Brokers:   []string{"localhost:9092"},
//		Topic:     "solar",
//		Partition: 0,
//		MinBytes:  10e3,
//		MaxBytes:  10e6,
//	})
//	defer r.Close()
//
//	for {
//		m, err := r.FetchMessage(context.Background())
//		if err != nil {
//			break
//		}
//
//		// 解析 Kafka 消息为请求类型
//		err = json.Unmarshal(m.Value, &req)
//		if err != nil {
//
//		}
//
//		// 将消息写入本地缓存以及发送到批处理系统
//		s.localCache.Set(req.UserID, req, cache.RedisContext)
//		s.batcher.Put(req.UserID, &KafkaData{
//			Key:   string(m.Key),
//			Value: req,
//		})
//	}
//}
