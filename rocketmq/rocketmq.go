package rocketmq

import (
	"errors"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"

	"github.com/qit-team/snow-core/config"
)

//依赖注入用的函数
func NewRocketMqClient(mqConfig config.RocketMqConfig) (client *RocketClient, err error) {
	// 初始化aliyunmq的 client
	defer func() {
		if e := recover(); e != nil {
			s := fmt.Sprintf("rocketmq client init panic: %v", e)
			err = errors.New(s)
		}
	}()

	if mqConfig.EndPoint != "" {
		client = new(RocketClient)
		//
		client.Consumer, err = rocketmq.NewPushConsumer(
			consumer.WithNameServer([]string{mqConfig.EndPoint}),
			consumer.WithCredentials(primitive.Credentials{
				AccessKey: mqConfig.AccessKey,
				SecretKey: mqConfig.SecretKey,
			}),
			consumer.WithGroupName(mqConfig.GroupId),
			consumer.WithNamespace(mqConfig.InstanceId),
			consumer.WithConsumerModel(consumer.Clustering),
			consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
			consumer.WithAutoCommit(true),
		)
		if err != nil {
			return nil, err
		}

		client.Producer, err = rocketmq.NewProducer(
			producer.WithNameServer([]string{mqConfig.EndPoint}),
			producer.WithCredentials(primitive.Credentials{
				AccessKey: mqConfig.AccessKey,
				SecretKey: mqConfig.SecretKey,
			}),
			producer.WithRetry(2),
			producer.WithGroupName(mqConfig.GroupId),
			producer.WithNamespace(mqConfig.InstanceId),
		)
		if err != nil {
			return nil, err
		}

	} else {
		err = errors.New("EndPoint empty, can not get client")
	}
	return
}

type RocketClient struct {
	Consumer rocketmq.PushConsumer
	Producer rocketmq.Producer
}
