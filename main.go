package main

import (
	"flag"
	"fmt"
	"zouyi/bluebell/controller"
	"zouyi/bluebell/dao/mysql"
	"zouyi/bluebell/dao/redis"
	"zouyi/bluebell/logger"
	"zouyi/bluebell/pkg/snowflake"
	"zouyi/bluebell/router"
	"zouyi/bluebell/setting"

	"go.uber.org/zap"
)

func main() {
	// --------flag-------------
	var configPath string
	flag.StringVar(&configPath, "conf", "D:/Codes/Golang/src/zouyi/bluebell/conf/config.yaml", "配置文件路径")
	flag.Parse()

	// 1. load config file
	if err := setting.Init(configPath); err != nil {
		fmt.Printf("init settings err: %s\n", err)
	}

	// 2. initialize log
	if err := logger.Init(setting.Conf.LogConfig, "dev"); err != nil {
		fmt.Printf("init logger err: %s\n", err)
	}
	defer zap.L().Sync()
	zap.L().Info("[logger] init success")

	startTime, machineId := setting.Conf.StartTime, setting.Conf.MachineID
	if err := snowflake.Init(startTime, machineId); err != nil {
		fmt.Printf("init snowflake id err: %v\n", err)
		return
	}
	zap.L().Info("[snowflake] init success")

	// 3. connect to database
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init logger err: %s\n", err)
	}
	defer mysql.Close()
	zap.L().Info("[mysql] init success")

	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init logger err: %s\n", err)
	}
	defer redis.Close()
	zap.L().Info("[redis] init success")

	// 4. route register
	if err := controller.InitTrans("en"); err != nil {
		zap.L().Error("init trans err", zap.Error(err))
		fmt.Printf("init validator err: %s\n", err)
		return
	}
	zap.L().Info("[trans] init success")

	r := router.Setup()
	zap.L().Info("[router] init success")

	//5. graceful shutdown
	//srv := &http.Server{
	//	Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
	//	Handler: r,
	//}
	//
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()
	//
	//quit := make(chan os.Signal, 1)
	//
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("server shutdown: ", err)
	//
	//}
	//zap.L().Info("server exiting")

	//// 5. start service
	r.Run(":8081")
}
