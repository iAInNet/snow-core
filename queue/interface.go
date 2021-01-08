package queue

import "context"

//队列驱动接口，所有队列驱动都需要实现以下接口
type Queue interface {
	//单入队
	Enqueue(ctx context.Context, key string, message string, args ...interface{}) (ok bool, err error)
	//单出队： 消息不存在是返回空字符串
	//增加返回参数，dequeueCount出队消费次数，目前只有alimns需要用到，redis暂无
	//增加入参，args ...interface{}
	Dequeue(ctx context.Context, key string, args ...interface{}) (message string, tag string, token string, dequeueCount int64, err error)
	//确认接收消息redis用不到，alimns需要，后续可以接入kafka或者rabbitmq
	// 增加入参args，因为rocketmq需要groupId等数据获取consumer
	AckMsg(ctx context.Context, key string, token string, args ...interface{}) (ok bool, err error)
	//单key批量入队
	BatchEnqueue(ctx context.Context, key string, messages []string, args ...interface{}) (ok bool, err error)
}
