package service

import (
	"context"
	"fmt"
	"github.com/ivpusic/grpool"
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/queue"
	"github.com/juju/ratelimit"
	"runtime"

)

type logicService struct {
	Pool *grpool.Pool
	LimitBucket *ratelimit.Bucket
	Object Object
}

func NewLogicService() *logicService {
	fn := func(buf []byte) {

		common.Cmd.Handle(buf)
		fmt.Println("a",string(buf))
	}
	var l = &logicService{Object: fn}
	l.Pool = grpool.NewPool(4, 32)

	//l.Object = l.Object
	return l
}

func (l *logicService) Start(ctx context.Context) {
	queue.Consumer(ctx,l.Go)
}

func (l *logicService) Close() {

}

//开启异步
func (l *logicService) Go(buf []byte) {
	l.Pool.JobQueue <- func() {
		l.Object(buf)
	}
}

type Object func(buf []byte)

func Recover(fn Object)  Object {
	return func(buf []byte)  {
		defer func() {
			if err := recover(); err != nil {
				stackBuf := make([]byte, 4096)
				n := runtime.Stack(stackBuf, false)
				golog.Error("Recover","err", err, "stackInfo", string(stackBuf[:n]))
			}
		}()
		fn(buf)
	}
}
