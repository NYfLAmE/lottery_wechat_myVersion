package repo

import (
	"github.com/sirupsen/logrus"
	"lottery_wechat/internal/model"
	"lottery_wechat/internal/pkg/gormcli"
)

func AddPrize(prizeList []*model.Prize) error {
	db := gormcli.GetDB() // 获取db连接

	if err := db.Model(&model.Prize{}).Create(prizeList).Error; err != nil { // 使用gorm的方法插入数据
		logrus.Errorf("repo AddPrize err: %v", err)
		return err
	}

	logrus.Infof("repo AddPrize success") // 打印成功日志
	return nil
}

func GetPrizeList() ([]*model.Prize, error) {
	db := gormcli.GetDB() // 获取db连接

	var prizeList []*model.Prize
	if err := db.Model(&model.Prize{}).Where("is_use = ?", 1).Find(&prizeList).Error; err != nil { // 使用gorm的方法查询数据
		logrus.Errorf("repo GetPrizeList err: %v", err)
		return nil, err
	}

	logrus.Infof("repo GetPrizeList success") // 打印成功日志
	return prizeList, nil
}
