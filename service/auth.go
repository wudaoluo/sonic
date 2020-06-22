package service

import (
	"database/sql"
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/dao"
	"github.com/wudaoluo/sonic/middleware"
	"github.com/wudaoluo/sonic/model"
)

type AuthService struct {
	
}

func NewAuthService() *AuthService{
	return &AuthService{}
}

func (a AuthService) Login(reqData *model.AuthLogin) (interface{},error) {
	imUser,err := dao.DBImUser.SelectByUserName(reqData.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			golog.Error("Login","func","dao.DBImUser.SelectByUserName","err",common.NOT_FOUND_ERROR)
			return nil, common.NOT_FOUND_ERROR
		}

		golog.Error("Login","func","dao.DBImUser.SelectByUserName","err",err)
		return nil, common.ErrEXec(err)
	}

	if imUser.Password != reqData.Password {
		golog.Error("Login","func","imUser.Password != reqData.Password","err",common.AUTH_ERR)
		return nil, common.AUTH_ERR
	}

	//jwt
	token,err := middleware.TokenGenerator(imUser)
	if err != nil {
		golog.Error("Login","func","middleware.TokenGenerator","err",err)
		return nil, common.SERVICE_ERROR
	}

	return token,nil
}

func (a AuthService) Logout(reqData interface{}) (interface{},error) {
	return nil, nil
}
