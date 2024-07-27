package api

import (
	"encoding/json"
	"fmt"
	"lottery_wechat/internal/model"
	"testing"
)

// 测试一下将奖品结构体 转换为 JSON串，方便后续使用这个JSON数据发送http请求来测试上传奖品的接口
func TestAddPrize(t *testing.T) {
	var prizeListHttp []*model.PrizeReq

	prizeListHttp = append(prizeListHttp, &model.PrizeReq{
		Name:  "奖品1",
		Type:  1,
		Total: 100,
	})

	prizeListHttp = append(prizeListHttp, &model.PrizeReq{
		Name:  "奖品2",
		Type:  2,
		Total: 100,
	})

	if prizeListJSON, err := json.Marshal(prizeListHttp); err != nil {
		t.Errorf("json.Marshal err: %v", err)
	} else {
		fmt.Println(string(prizeListJSON)) // 输出奖品列表结 对应的 JSON串
	}

}
