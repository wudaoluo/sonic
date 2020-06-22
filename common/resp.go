package common

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/sonic/model"
	"net/http"
)

func GinJsonResp(c *gin.Context,data interface{}) {
	c.JSON(http.StatusOK,model.GinResp{Data: data,Code: 0,Msg:""})
}

func GinJsonRespErr(c *gin.Context,err error) {
	c.JSON(http.StatusOK,model.GinResp{Data: nil,Code: 4000,Msg:err.Error()})
}