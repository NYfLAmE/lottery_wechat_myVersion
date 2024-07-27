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
