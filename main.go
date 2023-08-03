package main

import (
	"flag"
	"fmt"

	"github.com/YiZou89/bluebell/middleware/rabbitmq"

	"github.com/YiZou89/bluebell/controller"
	"github.com/YiZou89/bluebell/dao/mysql"
	"github.com/YiZou89/bluebell/dao/redis"
	"github.com/YiZou89/bluebell/logger"
	"github.com/YiZou89/bluebell/pkg/snowflake"
	"github.com/YiZou89/bluebell/router"
	"github.com/YiZou89/bluebell/setting"
	"go.uber.org/zap"
)

func main() {
	// 读取配置文件
	// ./bluebell -conf conf/config.yaml
	var configPath string
	flag.StringVar(&configPath, "conf", "./conf/config.yaml", "配置文件路径")
	flag.Parse()
	if err := setting.Init(configPath); err != nil {
		fmt.Printf("init settings err: %s\n", err)
	}

	// zap 日志记录
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger err: %s\n", err)
	}
	defer zap.L().Sync()

	// snowflake 雪花算法，用于注册时生成用于id
	startTime, machineId := setting.Conf.StartTime, setting.Conf.MachineID
	if err := snowflake.Init(startTime, machineId); err != nil {
		fmt.Printf("init snowflake id err: %v\n", err)
		return
	}

	// 连接 mysql 数据库
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql err: %s\n", err)
	}
	defer mysql.Close()

	// 连接 redis 数据库
	redis.Init(setting.Conf.RedisConfig)
	defer redis.Close()

	// 连接rabbitmq
	rabbitmq.Init(setting.Conf.RabbitmqConfig)
	defer rabbitmq.Close()

	// validator校验器
	if err := controller.InitTrans("en"); err != nil {
		fmt.Printf("init validator err: %s\n", err)
		return
	}

	// Gin 路由启动
	r := router.Setup(setting.Conf.Mode)

	//err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	//
	//if err != nil {
	//	fmt.Printf("run server failed, err:%v\n", err)
	//	return
	//}
	r.Run(":8081")
}
