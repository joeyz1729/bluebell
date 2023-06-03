package setting

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    uint16 `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
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
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
	//MinIdleConns int    `mapstructure:"min_idle_conns"`
}

var Conf = new(AppConfig)

//var ConfigFile = "D:/Codes/Golang/src/zouyi/web-app2/config/config.yaml"

func Init(configFile string) (err error) {
	viper.SetConfigFile(configFile)
	//viper.SetConfigName("./config.yaml")
	//viper.AddConfigPath("./config/") //在main.go中运行, 所以.是项目根目录
	//viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		//panic(fmt.Errorf("Fatal err read config file: %s\n", err))
		fmt.Println("[setting] viper read config file err ")
		return
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Println("[setting] viper unmarshal config file err")
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("[setting] viper config file has been changed")
		_ = viper.Unmarshal(&Conf)
	})
	return
}
