package queue

import(
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/common"
	"context"
	"github.com/wudaoluo/sonic/queue/queue_default"
)

type Queue interface {
	//Start()
	Processor
	Consumeor
}

type Processor interface {
	Producer(ctx context.Context,buf []byte) error
	ProducerClose() error
}

type Consumeor interface {
	Consumer(ctx context.Context,fn func(buf []byte))
	ConsumerClose() error
}


var queue Queue

func Producer(ctx context.Context,buf []byte) error {
	return queue.Producer(ctx,buf)
}


func ProducerClose() error {
	return queue.ProducerClose()
}

func Consumer(ctx context.Context,fn func(buf []byte)) {
	queue.Consumer(ctx,fn)
}

func ConsumerClose() error {
	return queue.ConsumerClose()
}

func Init()  {
	conf := &common.GetConf().Queue
	golog.Info("queue.init","type",conf.Type)
	switch conf.Type {
	case "kafka":
		//queue = queue_kafka.New(conf)
	case "default":
		queue = queue_default.New(conf)
	default:
		golog.Error("queue.init","err",common.NOT_FOUND_ERROR)
		panic(common.NOT_FOUND_ERROR)
	}
}