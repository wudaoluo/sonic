package middleware

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/model"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"errors"
)
/*
Playload(载荷又称为Claim)
iss: 签发者
sub: 面向的用户
aud: 接收方
exp: 过期时间
nbf: 生效时间
iat: 签发时间
jti: 唯一身份标识

HMACSHA256(
    base64UrlEncode(header) + "." +
    base64UrlEncode(payload),
    secret
)
*/

const (
	TOKEN = "Token"
	UID = "uid"
	USERNAME = "username"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		user,err := parseToken(c.GetHeader(TOKEN))
		if err != nil {
			golog.Error("middleware.jwt","func","parseToken","err",err)
			return
		}
		c.Set(UID,user.Uid)
		c.Set(USERNAME,user.Username)
	}
}

func secret()jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(common.GetConf().Auth.Jwt.Key),nil
	}
}


func parseToken(tokenString string)(*model.ImUser,error){
	user := &model.ImUser{}
	token,err := jwt.Parse(tokenString,secret())
	if err != nil{
		return nil,err
	}
	claim,ok := token.Claims.(jwt.MapClaims)
	if !ok{
		err = errors.New("cannot convert claim to mapclaim")
		return nil,err
	}
	//验证token，如果token被修改过则为false
	if  !token.Valid{
		err = errors.New("token is invalid")
		return nil,err
	}

	if v,ok := claim[UID];ok {
		user.Uid = cast.ToInt64(v)
	}

	if v,ok := claim[USERNAME];ok {
		user.Username = cast.ToString(v)
	}

	return user,nil
}

func TokenGenerator(user *model.ImUser) (string, error) {
	jwtConf := common.GetConf().Auth.Jwt
	now := time.Now()
	expire := now.Add(jwtConf.Timeout*time.Second)

	claim := jwt.MapClaims{
		UID:       user.Uid,
		USERNAME: user.Username,
		"exp":     expire.Unix(),
		"iat":     now,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	return token.SignedString([]byte(jwtConf.Key))
}
