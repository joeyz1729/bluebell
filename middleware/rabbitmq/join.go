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
	RmqJoinAdd  *JoinChan
	RmqJoinDel  *JoinChan
	addJoinName = "join_add"
	delJoinName = "join_del"
)

type JoinChan struct {
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

func NewJoinRmq(queueName string) (*JoinChan, error) {
	joinMq := &JoinChan{
		queueName: queueName,
	}
	channel, err := RabbitmqConn.Channel()
	if err != nil {
		zap.L().Error("failed to open join channel", zap.Error(err))
		return nil, err
	}
	joinMq.channel = channel
	return joinMq, nil

}

func (j *JoinChan) Close() {
	_ = j.channel.Close()
}

func (j *JoinChan) Publish(message string) error {
	_, err := j.channel.QueueDeclare(
		j.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("declare queue err: ", zap.Error(err))
		return err
	}
	err = j.channel.Publish(
		j.exchange,
		j.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		zap.L().Error("publish message to JoinChan err: ", zap.Error(err))
		return err
	}

	zap.L().Debug("publish message to join channel success")
	return nil
}

func (j *JoinChan) Consumer() {
	_, err := j.channel.QueueDeclare(
		j.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("declare join queue err:", zap.Error(err))
		return
	}

	msgs, err := j.channel.Consume(
		j.queueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		zap.L().Error("join queue consume err:", zap.Error(err))
		return
	}
	var forever chan struct{}
	switch j.queueName {
	case addJoinName:
		go j.consumerFollowAdd(msgs)
	case delJoinName:
		go j.consumerFollowDel(msgs)
	}
	zap.L().Debug("[*] waiting for messages, to exit press Ctrl+C")
	<-forever
}

func (j *JoinChan) consumerFollowAdd(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		uid, _ := strconv.Atoi(params[0])
		toUid, _ := strconv.Atoi(params[1])
		zap.L().Debug("[.] rabbitmq consumer, start add join relations",
			zap.Int("uid", uid),
			zap.Int("toUid", toUid),
		)

		if err := mysql.AddFollow(uid, toUid); err != nil {
			zap.L().Error("insert join relation err", zap.Error(err))
			return

		}
	}

}

func (j *JoinChan) consumerFollowDel(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		uid, _ := strconv.Atoi(params[0])
		toUid, _ := strconv.Atoi(params[1])
		zap.L().Debug("[.] rabbitmq consumer, start add join relations",
			zap.Int("uid", uid),
			zap.Int("toUid", toUid),
		)
		if err := mysql.AddFollow(uid, toUid); err != nil {
			zap.L().Error("insert join relation err: ", zap.Error(err))
			return
		}
		zap.L().Info("insert join relation success")

	}
}

func InitJoinRmq() (err error) {
	RmqJoinAdd, err = NewJoinRmq(addQueueName)
	if err != nil {
		zap.L().Error("init join add err", zap.Error(err))
		return
	}
	RmqJoinDel, err = NewJoinRmq(delQueueName)
	if err != nil {
		zap.L().Error("init join del err", zap.Error(err))
		return
	}
	go RmqFollowAdd.Consumer()
	go RmqFollowDel.Consumer()
	zap.L().Info("init follow rmq success")
	return nil
}
