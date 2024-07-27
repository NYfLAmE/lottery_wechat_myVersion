package main

import (
	"fmt"
	"lottery_wechat/configs"
	"lottery_wechat/router"
)

func Init() { // 专门设置一个init函数，用于将程序执行真正的业务逻辑前对相关环境以及配置初始化
	configs.InitLogrus() // 目前只需要初始化一个logrus 20240725
}

func main() {
	config := configs.GetGlobalConfig()

	fmt.Println(config)

	Init()

	r := router.SetRouter()
	if err := r.Run(fmt.Sprintf(":%d", config.AppConfig.Port)); err != nil {
		panic(fmt.Sprintf("route启动失败：%v", err))
	}
}
