package queue_default

import (
	"github.com/wudaoluo/sonic/model"
	"sync/atomic"
	"context"
)

type QueueDefault struct {
	queue chan []byte
	status int32
}

func New(conf *model.ConfigQueue) *QueueDefault {
	return &QueueDefault{
		queue:make(chan []byte, 100),
		status: 1,
	}
}

func (q *QueueDefault) Consumer(ctx context.Context,fn func(buf []byte)) {
	go func() {
		for msg := range q.queue {
			fn(msg)
		}
	}()

}

func (q *QueueDefault) ProducerClose() error {
	if atomic.LoadInt32(&q.status) == 1 {
		atomic.StoreInt32(&q.status,0)
		close(q.queue)
	}

	return nil
}

func (q *QueueDefault) ConsumerClose() error {
	if atomic.LoadInt32(&q.status) == 1 {
		atomic.StoreInt32(&q.status,0)
		close(q.queue)
	}


	return nil
}

func (q *QueueDefault) Producer(ctx context.Context,buf []byte) error {
	if atomic.LoadInt32(&q.status) == 1 {
		q.queue <- buf
	}

	return nil
}