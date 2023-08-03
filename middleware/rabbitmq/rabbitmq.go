package rabbitmq

import (
	"fmt"

	"github.com/YiZou89/bluebell/setting"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var RabbitmqConn *amqp.Connection

//const MQURL = "amqp://tiktok:tiktok@106.14.75.229:5672/"

var RabbitmqURL string

func Init(conf *setting.RabbitmqConfig) (err error) {
	RabbitmqURL = fmt.Sprintf("amqp://%s:%s@%s:%d/",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
	)
	dial, err := amqp.Dial(RabbitmqURL)
	if err != nil {
		zap.L().Error("connect to rabbitmq err: ", zap.Error(err))
		return
	}
	zap.L().Info("[rabbitmq] init success")
	RabbitmqConn = dial
	return nil
}
func Close() {
	RabbitmqConn.Close()
}
