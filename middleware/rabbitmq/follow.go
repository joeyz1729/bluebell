package rabbitmq

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/YiZou89/bluebell/dao/mysql"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var (
	RmqFollowAdd *FollowChan
	RmqFollowDel *FollowChan
	addQueueName string = "follow_add"
	delQueueName string = "follow_del"
)

type FollowChan struct {
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

func NewFollowRmq(queueName string) (*FollowChan, error) {
	followMq := &FollowChan{
		queueName: queueName,
	}
	channel, err := RabbitmqConn.Channel()
	if err != nil {
		zap.L().Error("failed to open follow channel", zap.Error(err))
		return nil, err
	}
	followMq.channel = channel
	return followMq, nil
}

func (f *FollowChan) Close() {
	_ = f.channel.Close()
}

func (f *FollowChan) Publish(message string) error {
	// 声明队列
	_, err := f.channel.QueueDeclare(
		f.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		zap.L().Error(fmt.Sprintf("declare queue %s err", f.queueName), zap.Error(err))
		return err
	}

	// 发送消息
	err = f.channel.Publish(
		f.exchange,
		f.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		zap.L().Error("publish err", zap.Error(err))
		return err
	}
	zap.L().Debug("publish success")
	return nil
}

func (f *FollowChan) Consumer() {
	_, err := f.channel.QueueDeclare(
		f.queueName,
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		zap.L().Error(fmt.Sprintf("declare queue %s err", f.queueName), zap.Error(err))
		//return err
	}

	// consume
	msgs, err := f.channel.Consume(
		f.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var forever chan struct{}
	switch f.queueName {
	case addQueueName:
		go f.consumerFollowAdd(msgs)
	case delQueueName:
		go f.consumerFollowDel(msgs)
	}
	zap.L().Debug("[*] waiting for messages, to exit press Ctrl+C")
	<-forever
}

func (f *FollowChan) consumerFollowAdd(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		uid, _ := strconv.Atoi(params[0])
		toUid, _ := strconv.Atoi(params[1])
		zap.L().Debug("[.] rabbitmq consumer, start add follow relations",
			zap.Int("uid", uid),
			zap.Int("toUid", toUid),
		)
		// 开始添加到数据库
		if err := mysql.AddFollow(uid, toUid); err != nil {
			zap.L().Error("insert follow relation err", zap.Error(err))
			return
		}
		zap.L().Error("insert follow relation success")

	}
}

func (f *FollowChan) consumerFollowDel(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		uid, _ := strconv.Atoi(params[0])
		toUid, _ := strconv.Atoi(params[1])
		zap.L().Debug("[.] rabbitmq consumer, start add follow relations",
			zap.Int("uid", uid),
			zap.Int("toUid", toUid),
		)
		// 开始添加到数据库
		if err := mysql.AddFollow(uid, toUid); err != nil {
			zap.L().Error("insert follow relation err", zap.Error(err))
			return
		}
		zap.L().Info("insert follow relation success")

	}
}

func InitFollowRmq() (err error) {
	RmqFollowAdd, err = NewFollowRmq(addQueueName)
	if err != nil {
		zap.L().Error("init follow add rmq err", zap.Error(err))
		return
	}
	go RmqFollowAdd.Consumer()

	RmqFollowDel, err = NewFollowRmq(delQueueName)
	if err != nil {
		zap.L().Error("init follow del rmq err", zap.Error(err))
		return
	}
	go RmqFollowDel.Consumer()
	zap.L().Info("init follow rmq success")
	return nil
}
