package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/middleware"
	"github.com/wudaoluo/sonic/queue"
)

func LogicV1Router(parentRoute gin.IRouter) {
	router := parentRoute.Group("/logic")
	router.Use(middleware.Jwt())
	end := NewLogic()
	router.POST("/command",end.Command)
}

type logic struct {
}

func NewLogic() *logic {
	return &logic{
	}
}

//用来模拟的接口
func (l logic) Command(c *gin.Context) {
	buf,err := c.GetRawData()
	if err != nil {
		common.GinJsonRespErr(c,common.PARAM_ERROR)
		return
	}
	err = queue.Producer(c,buf)
	if err != nil {
		common.GinJsonRespErr(c,common.SERVICE_ERROR)
		return
	}

	//ret,err := common.Cmd.Handle(&common.CmdContext{CmdType: common.MsgReceive, Args: " Post"})
	//if err != nil {
	//	common.GinJsonRespErr(c,err)
	//	return
	//}
	common.GinJsonResp(c,true)
}

