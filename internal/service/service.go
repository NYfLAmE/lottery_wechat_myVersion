package service

import (
	"github.com/sirupsen/logrus"
	"lottery_wechat/internal/model"
	"lottery_wechat/internal/repo"
)

func AddPrize(prizeListHttp []*model.PrizeReq) error { // 将http中的奖品列表中的奖品 映射到 数据库表中
	// 定义一个全局数据库prizeList表变量，用来暂存遍历过程中的prizeDB
	var prizeListDB []*model.Prize

	// 1. 遍历http奖品列表，将其中的每个奖品 都映射到 数据库表中
	for _, prizeHttp := range prizeListHttp {
		// 首先将prizeHttp 转换成 数据库中的奖品类型
		prizeDB := model.Prize{
			Name:  prizeHttp.Name,
			Type:  prizeHttp.Type,
			Total: prizeHttp.Total,
		}

		// 2. 将prizeDB 暂存到到prizeListDB中
		prizeListDB = append(prizeListDB, &prizeDB)
	}

	// 3. 调用repo层的方法 将prizeListDB 插入到数据库中
	if err := repo.AddPrize(prizeListDB); err != nil {
		logrus.Errorf("service AddPrize err: %v", err)
		return err
	}

	return nil
}

func GetPrizeInfo() (*model.GetPrizeRsp, error) {

	// 定义一个全局变量 prizeListDB，存放repo层返回到记录集合
	var prizeListDB []*model.Prize

	// 调用repo层的方法，从数据库中检索出所有的记录
	prizeListDB, err := repo.GetPrizeList()
	if err != nil {
		logrus.Errorf("service GetPrizeInfo err: %v", err)
		return nil, err
	}

	// 处理数据，将结果返回给api层
	var typeNum int64 = 0
	var total int64 = 0
	for _, prizeDB := range prizeListDB {
		typeNum++
		total += prizeDB.Total
	}

	// 返回处理结果
	return &model.GetPrizeRsp{
		TypeNum: typeNum,
		Total:   total,
	}, nil

}
