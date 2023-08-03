package setting

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name                       string        `mapstructure:"name"`
	Mode                       string        `mapstructure:"mode"`
	Version                    string        `mapstructure:"version"`
	Host                       string        `mapstructure:"host"`
	Port                       int           `mapstructure:"port"`
	StartTime                  string        `mapstructure:"start_time"`
	MachineID                  uint16        `mapstructure:"machine_id"`
	AccessTokenExpireDuration  time.Duration `mapstructure:"jwt_access_expire"`
	RefreshTokenExpireDuration time.Duration `mapstructure:"jwt_refresh_expire"`
	*RabbitmqConfig            `mapstructure:"rabbitmq"`
	*LogConfig                 `mapstructure:"log"`
	*MySQLConfig               `mapstructure:"mysql"`
	*RedisConfig               `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	Username     string        `mapstructure:"username"`
	Password     string        `mapstructure:"password"`
	DB           string        `mapstructure:"dbname"`
	MaxLifetime  time.Duration `mapstructure:"max_life_time"`
	MaxOpenConns int           `mapstructure:"max_open_conns"`
	MaxIdleConns int           `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
	//MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type RabbitmqConfig struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	ManagementPort int    `mapstructure:"management_port"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
}

var Conf = new(AppConfig)

func Init(configFile string) (err error) {
	// 读取配置文件
	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		//panic(fmt.Errorf("Fatal err read config file: %s\n", err))
		fmt.Printf("[setting] viper read config file err:%v\n", err)
		return
	}
	// 解析配置文件
	err = viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Printf("[setting] viper unmarshal config file err:%v\n", err)
		return
	}

	// 实时监控配置变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("[setting] viper config file has been changed")
		_ = viper.Unmarshal(&Conf)
	})
	return
}
