package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/common"
	"github.com/wudaoluo/sonic/model"
	"time"
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
ref: 刷新token

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
	REFRESH = "ref"
)

type JwtToken struct {
	Token string `json:"token"`
	TokenRefresh string `json:"token_refresh"`
}

func (j *JwtToken) Gen(user *model.ImUser) error {
	tokenStr,err := TokenGenerator(user,false)
	if err != nil {
		return err
	}

	j.Token = tokenStr

	tokenRefreshStr,err := TokenGenerator(user,true)
	if err != nil {
		return err
	}
    j.TokenRefresh = tokenRefreshStr
	return nil
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		user,err := ParseToken(c.GetHeader(TOKEN),false)
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


func ParseToken(tokenString string,refresh bool)(*model.ImUser,error){
	token ,err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claim,ok := token.Claims.(jwt.MapClaims)
	if !ok{
		err = errors.New("cannot convert claim to mapclaim")
		return nil,err
	}

	user := &model.ImUser{}
	if refresh {
		if v,ok := claim[REFRESH];ok && !cast.ToBool(v){
			return nil, common.TOKEN_INVALID
		}
	}

	if v,ok := claim[UID];ok {
		user.Uid = cast.ToInt64(v)
	}

	if v,ok := claim[USERNAME];ok {
		user.Username = cast.ToString(v)
	}

	return user,nil
}

func VerifyToken(tokenString string) (*jwt.Token,error){
	token,err := jwt.Parse(tokenString,secret())
	if err != nil{
		golog.Error("VerifyToken","func","jwt.Parse","token",tokenString,"err",err)
		return nil,err
	}

	//验证token，如果token被修改过则为false
	if  !token.Valid{
		err = common.TOKEN_INVALID
		golog.Error("VerifyToken","func","token.Valid","token",tokenString,"err",err)
		return nil,err
	}

	return token, nil
}


func TokenGenerator(user *model.ImUser,refresh bool) (string, error) {
	jwtConf := common.GetConf().Auth.Jwt
	now := time.Now()

	expire := now.Add(jwtConf.Timeout*time.Second)
	if refresh {
		expire = now.Add(2*jwtConf.Timeout*time.Second)
	}
	claim := jwt.MapClaims{
		UID:       user.Uid,
		USERNAME:  user.Username,
		"exp":     expire.Unix(),
		"iat":     now,
		REFRESH:   refresh,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	return token.SignedString([]byte(jwtConf.Key))
}
