package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/middleware"
	"github.com/wudaoluo/sonic/model"
	"github.com/wudaoluo/sonic/service"
)

func LogicV1Router(parentRoute gin.IRouter) {
	router := parentRoute.Group("/auth")
	end := NewAuth()
	router.POST("/login",middleware.Jwt(),end.Login)
	router.POST("/logout",end.Logout)
	router.POST("/token/refresh",end.TokenRefresh)
}

type logic struct {
	service *service.AuthService
}

func NewLogic() *auth {
	return &auth{
		service.NewAuthService(),
	}
}

func (l logic) Command(c *gin.Context) {
	var req model.AuthLogin
	if err := c.Bind(&req); err != nil {
		common.GinJsonRespErr(c,common.PARAM_ERROR)
		return
	}

	ret,err := common.Cmd.Handle(&common.CmdContext{CmdType: common.MsgReceive, Args: " Post"})
	if err != nil {
		common.GinJsonRespErr(c,err)
		return
	}
	common.GinJsonResp(c,ret)
}

