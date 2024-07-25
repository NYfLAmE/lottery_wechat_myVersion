package configs

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

var (
	globalConfig GlobalConfig // 创建一个全局配置变量
	once         sync.Once
)

type GlobalConfig struct { // 创建一个大的配置类型把所有的配置类型都包含进来
	AppConfig AppConf `yaml:"app" mapstructure:"app"`
	LogConfig LogConf `yaml:"log" mapstructure:"log"`
}

type AppConf struct { // 定义一个App的配置类型
	AppName string `yaml:"app_name" mapstructure:"app_name"`
	Version string `yaml:"version" mapstructure:"version"`
	Port    int    `yaml:"port" mapstructure:"port"`
	RunMode string `yaml:"run_mode" mapstructure:"run_mode"`
}

type LogConf struct { // 定义一个Log的配置类型
	LogPattern string `yaml:"log_pattern" mapstructure:"log_pattern"`
	LogPath    string `yaml:"log_path" mapstructure:"log_path"`
	SaveDays   int    `yaml:"save_days" mapstructure:"save_days"`
	Level      string `yaml:"level" mapstructure:"level"`
}

// 成功引入viper包之后，使用viper库来读取配置文件
func readConf() {
	viper.SetConfigName("config")    // 设置配置文件名称
	viper.SetConfigType("yaml")      // 设置配置文件类型
	viper.AddConfigPath("./configs") // 设置配置文件路径，.是根目录，也就是程序启动的路径，所以配置文件的路径为./configs

	// 通过上面三个操作，配置文件已经加载到viper中了，接下来就是使用viper来读取配置文件
	if err := viper.ReadInConfig(); err != nil { // 读取配置文件
		panic("read config file error" + err.Error()) // 读取配置文件失败，panic，并将错误消息给打印出来
	}

	// 读取配置文件成功，接下来就是把配置文件中的内容给解析到结构体中
	if err := viper.Unmarshal(&globalConfig); err != nil { // 将配置文件中的内容映射到globalConfig这个结构体里面
		panic("unmarshal config file error" + err.Error()) // 解析配置文件失败，panic，并将错误消息给打印出来
	}

}

func GetGlobalConfig() *GlobalConfig { // 定义一个get函数，方便外部使用配置
	once.Do(readConf)    // 只执行一次，也就是只初始化globalConfig这个结构体变量一次
	return &globalConfig // 每次调用GetGlobalConfig，返回的都是同一个globalConfig
}

func InitLogrus() {
	// 我们的日志是通过第三方包logrus来完成的
	// 我们需要将yml中的配置【可以通过getGlobalConfig得到】初始化到logrus里面以应用到日志的生成与输出
	// 我们这里写一个InitLogrus函数来完成这个操作

	logrus.SetFormatter(&logFormatter{ // 将自定义的logFormatter作为logrus的Formatter
		logrus.TextFormatter{ // 初始化自定义的logFormatter的成员TextFormatter
			DisableColors:   false,                 // 是否禁用颜色
			FullTimestamp:   true,                  // 是否显示完整的时间戳
			TimestampFormat: "2006-01-02 15:04:05", // 时间戳的格式
		},
	})

	logrus.SetReportCaller(true) // 设置是否显示调用者的信息，这里为true，后续Format要格式化的Entry.hasCaller就为true，也就可以打印调用信息了

	config := GetGlobalConfig() // 获取yml中的配置信息 以 初始化logrus

	// 初始化logrus的日志级别
	level, err := logrus.ParseLevel(config.LogConfig.Level) // 将全局配置信息中对level信息解析成logrus能够使用的日志级别（也就是logrus.Level类型的数据）
	if err != nil {                                         // 如果解析失败，那么就panic
		panic("parse log level error" + err.Error())
	}

	logrus.SetLevel(level) // 设置logrus的日志级别

	// 接下来就是初始化logrus的输出方式，也就是将日志写入到哪里，同样，也是从yml配置信息config实例里面取
	switch config.LogConfig.LogPattern {
	case "stdout": // 如果是stdout，那么就将日志写入到标准输出中
		logrus.SetOutput(os.Stdout)
	case "stderr": // 如果是stderr，那么就将日志写入到标准错误输出中
		logrus.SetOutput(os.Stderr)
	case "file": // 如果是file，那么就将日志写入到文件中，我们使用第三方包来自定义这个文件的相关属性
		logRotate, err := rotatelogs.New( // 得到一个rotatelogs包下的一个专门用于记录日志的文件，这个文件支持根据下面设定的WithRotation...参数
			// 实现automatically rotated，也就是自动帮我们管理生成的日志文件
			config.LogConfig.LogPath+".%Y%m%d%H%M%S",                                     // 定义日志文件的名称
			rotatelogs.WithMaxAge(time.Duration(config.LogConfig.SaveDays)*24*time.Hour), // 定义日志文件的最大保存时间
		)
		if err != nil {
			panic("log rotate conf error" + err.Error())
		}
		logrus.SetOutput(logRotate) // 设置logrus的输出方式，输出到logRotate这个日志文件中
	default: // 如果是其他的，那么就panic
		panic("unsupported log pattern" + config.LogConfig.LogPattern)
	}
}
