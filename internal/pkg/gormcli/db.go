package gormcli

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lottery_wechat/configs"
	"sync"
	"time"
)

var db *gorm.DB    // 定义一个全局的db变量，方便外部获取以及使用，跟获取globalConfig一样，要做单例模式
var once sync.Once // 用于实现db单例，只初始化一次

func OpenDB() { // 定义一个OpenDB函数，用于打开db（一个数据库连接），并在db上做一些连接配置
	dbConfig := configs.GetGlobalConfig().DBConfig // 获取db配置

	connArg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName) // 拼接连接数据库的参数

	logrus.Infof("connectArg:%s\n", connArg) // 打印一下日志

	// 连接数据库
	var err error
	db, err = gorm.Open(mysql.Open(connArg), &gorm.Config{}) // 注意：在gorm.Open函数中使用了mysql.Open函数，
	// mysql这个包与需要提前引入：gorm.io/driver/mysql
	if err != nil {
		panic("failed to connect database")
	}

	// 设置数据库连接池的参数
	sqlDB, err := db.DB() // db.DB()会返回一个表示当前数据库连接池的对象
	if err != nil {
		panic("failed to get sqlDB")
	}

	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)                                // 设置数据库连接池的最大空闲连接数（指长连接)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)                                // 设置数据库连接池的最大打开连接数（包括长连接短连接）
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifetime) * time.Second) // 设置连接池中一个连接的最大存活时间
}

func GetDB() *gorm.DB { // 定义一个GetDB函数，用于获取数据库连接
	once.Do(OpenDB) // 调用OpenDB函数，但是只执行一次
	return db       // 返回db变量
}
