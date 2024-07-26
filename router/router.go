package router

import (
	"github.com/gin-gonic/gin"
	"lottery_wechat/api"
)

func SetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // 生产环境
	r := gin.Default()

	group := r.Group("/lottery_wechat")
	group.POST("/add_prize", api.AddPrize)

	return r
}
