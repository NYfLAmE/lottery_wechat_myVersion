package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lottery_wechat/internal/model"
	"lottery_wechat/internal/service"
	"net/http"
)

func AddPrize(ctx *gin.Context) { // 负责解析/add_prize请求发来的内容 & 调用service层的处理方法
	req := model.AddPrizeReq{} // 获取一个AddPrizeReq类型的变量，用来存放请求中的数据

	if err := ctx.BindJSON(&req); err != nil { // 解析请求中的数据
		logrus.Errorf("/add_prize httpReq parse err: %v", err) // 打印日志
		ctx.JSON(http.StatusBadRequest, 200)                   // 请求中的数据不符合我们的要求，解析失败，返回一个badRequest的响应
		return
	}

	// 参数拿出来后，如果业务需要做参数校验，可以在这里做

	// 调用service 添加奖品
	if err := service.AddPrize(req.PrizeList); err != nil { // 调用service层的处理方法
		logrus.Errorf("api AddPrize err: %v", err)    // 打印日志
		ctx.JSON(http.StatusInternalServerError, 500) // 返回错误码
		return
	}

	ctx.JSON(http.StatusOK, "AddPrize Success") // 成功返回一个ok的响应
}
