package common

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	//EXPIRED_ERROR = engine.NewResponseError(ERROR_STATUS_UNAUTHORIZED, "token expired")
	//AD_USERINFO_ERROR = engine.NewResponseError(ERROR_LOGIQUE_FAILURE, "ad userInfo is empty")
	// 参数错误
	PARAM_ERROR     = errors.New("param is error")
	NOT_FOUND_ERROR = errors.New("not found")
	SERVICE_ERROR   = errors.New("服务器错误")
	AUTH_ERROR      = errors.New("认证失败")
	TOKEN_INVALID   = errors.New("token is invalid")
	COMMAND_INVALID = errors.New("invalid Command")
	COMMAND_NIL     = errors.New("command is nil")

	DB_NOT_FOUND_ERR = errors.New("db item not found")
	DB_INSERT_ERR    = errors.New("db insert faild")
	DB_PARAM_ERROR   = errors.New("db param is error")
	//PARAM_ERROR = engine.NewResponseError(ERROR_PARAM_ERROR, "param is error")
	//MYSQL_PARAM_ERROR = engine.NewResponseError(ERROR_PARAM_ERROR, "sql param is error")
)

func ErrEXec(err error) error {
	if gin.Mode() != gin.ReleaseMode {
		return err
	}

	return SERVICE_ERROR
}
