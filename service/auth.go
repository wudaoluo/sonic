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
		golog.Error("Login","func","imUser.Password != reqData.Password","err",common.AUTH_ERROR)
		return nil, common.AUTH_ERROR
	}

	//jwt
	jwtToken := middleware.JwtToken{}
	err = jwtToken.Gen(imUser)
	if err != nil {
		golog.Error("Login","func","jwtToken.Gen","err",err)
		return nil, common.SERVICE_ERROR
	}

	return jwtToken,nil
}

func (a AuthService) Logout(reqData interface{}) (interface{},error) {
	return nil, nil
}

func (a AuthService) TokenRefresh(reqData *model.AuthTokenRefresh) (interface{},error) {
	imUser,err := middleware.ParseToken(reqData.TokenRefresh,true)
	if err != nil {
		golog.Error("TokenRefresh","func","middleware.ParseToken","err",err)
		return nil, err
	}
	jwtToken := middleware.JwtToken{}
	err = jwtToken.Gen(imUser)
	if err != nil {
		golog.Error("TokenRefresh","func","jwtToken.Gen","err",err)
		return nil, common.SERVICE_ERROR
	}

	return jwtToken,nil

}
