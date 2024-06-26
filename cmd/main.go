package main

import (
	"Solar_motion/config"
	"Solar_motion/pkg/utils/log"
	"Solar_motion/repository/cache"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/es"
	"Solar_motion/routers"
	"fmt"
)

func main() {
	loading() // 加载配置
	r := routers.NewRouter()
	_ = r.Run(config.Config.System.HttpPort)
	fmt.Println("启动配成功...")
}

// loading一些配置
func loading() {
	config.InitConfig()
	dao.InitMySQL()
	cache.InitCache()
	log.InitLog()
	es.Es()
	fmt.Println("加载配置完成...")
}
