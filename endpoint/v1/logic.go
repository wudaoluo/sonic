package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/middleware"
	"github.com/wudaoluo/sonic/model"
	"github.com/wudaoluo/sonic/queue"
	"github.com/wudaoluo/sonic/service"
)

func LogicV1Router(parentRoute gin.IRouter) {
	router := parentRoute.Group("/logic")
	router.Use(middleware.Jwt())
	end := NewLogic()
	router.POST("/command", end.Command)
	router.POST("/msg/receive", end.MsgReceive)
}

type logic struct {
}

func NewLogic() *logic {
	return &logic{}
}

//用来模拟的接口
func (l logic) Command(c *gin.Context) {
	buf, err := c.GetRawData()
	if err != nil {
		common.GinJsonRespErr(c, common.PARAM_ERROR)
		return
	}
	err = queue.Producer(c, buf)
	if err != nil {
		common.GinJsonRespErr(c, common.SERVICE_ERROR)
		return
	}

	common.GinJsonResp(c, true)
}

//接受消息
func (l logic) MsgReceive(c *gin.Context) {
	var param model.LogicCommand
	if err := c.Bind(&param); err != nil {
		common.GinJsonRespErr(c, common.PARAM_ERROR)
		return
	}

	msg := service.MsgReceive{}
	ret, err := msg.Do(param.Data)
	if err != nil {
		common.GinJsonRespErr(c, common.SERVICE_ERROR)
		return
	}

	common.GinJsonResp(c, ret)
}
