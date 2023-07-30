package mysql

import (
	"fmt"
	"time"

	"github.com/YiZou89/bluebell/setting"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(conf *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DB,
	)
	//"root:mysql@tcp(127.0.0.1:3306)/bluebell?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect to mysql failed, err: %v\n", zap.Error(err))
		return
	}
	err = db.Ping()
	if err != nil {
		zap.L().Error("ping mysql failed, err: %v\n", zap.Error(err))
		return
	}

	db.SetConnMaxLifetime(conf.MaxLifetime * time.Second)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)

	zap.L().Info("[mysql] init success")
	return nil
}

func Close() {
	_ = db.Close()
}
