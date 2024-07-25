package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"lottery_wechat/configs"
	"time"
)

func Init() { // 专门设置一个init函数，用于将程序执行真正的业务逻辑前对相关环境以及配置初始化
	configs.InitLogrus() // 目前只需要初始化一个logrus 20240725
}

func main() {
	config := configs.GetGlobalConfig()

	fmt.Println(config)

	Init()
	for i := 1; ; i++ {
		if i%2 == 0 {
			logrus.Infof("这是第%d次打印日志", i)
		} else {
			logrus.Errorf("这是第%d次打印日志", i)
		}

		time.Sleep(time.Second)
	}
}
