package es

import (
	"Solar_motion/config"
	"context"
	"github.com/olivere/elastic/v7"
)

// //，Elasticsearch 被用作一个日志和事件数据的存储系统
//
// var EsClient *elastic.Client
//
// // InitEs 初始化es
//
//	func InitEs() {
//		eConfig := config.Config.Es
//		esConn := fmt.Sprintf("http://%s:%s", eConfig.EsHost, eConfig.EsPort)
//		cfg := elastic.Config{
//			Addresses: []string{esConn},
//		}
//		client, err := elastic.NewClient(cfg)
//		if err != nil {
//			log.Panic(err)
//		}
//		EsClient = client
//	}
//
// // EsHookLog 初始化log日志
//
//	func EsHookLog() *eslogrus.ElasticHook {
//		eConfig := config.Config.Es
//		hook, err := eslogrus.NewElasticHook(EsClient, eConfig.EsHost, logrus.DebugLevel, eConfig.EsIndex)
//		if err != nil {
//			log.Panic(err)
//		}
//		return hook
//	}
var (
	esClient *elastic.Client
)

func Es() {
	rConfig := config.Config.Es
	client, _ := elastic.NewClient(elastic.SetURL(rConfig.URL), elastic.SetSniff(false))
	if _, _, err := client.Ping(rConfig.URL).Do(context.Background()); err != nil {
		panic(err)
	}
	esClient = client
}

func NewEs() *elastic.Client {
	return esClient
}
