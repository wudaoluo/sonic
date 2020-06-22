package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/middleware"
	"github.com/wudaoluo/sonic/model"
	"github.com/wudaoluo/sonic/service"
)

func AuthV1Router(parentRoute gin.IRouter) {
	router := parentRoute.Group("/auth")
	end := NewAuth()
	router.POST("/login",middleware.Jwt(),end.Login)
	router.POST("/logout",end.Logout)
}

type auth struct {
	service *service.AuthService
}

func NewAuth() *auth {
	return &auth{
		service.NewAuthService(),
	}
}

func (a *auth) Login(c *gin.Context) {
	var req model.AuthLogin
	if err := c.Bind(&req); err != nil {
		common.GinJsonRespErr(c,common.PARAM_ERROR)
		return
	}

	ret,err := a.service.Login(&req)
	if err != nil {
		common.GinJsonRespErr(c,err)
		return
	}
	common.GinJsonResp(c,ret)
}

func (a auth) Logout(c *gin.Context) {

}