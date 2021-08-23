package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/utils"
	"star-server/utils/errmsg"
	"strings"
	"time"
)

var JwyKey = []byte(utils.JwtKey)

type MyClaims struct {
	Openid string `json:"openid"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(openid string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaim := MyClaims{
		openid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jimyag",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaim)
	token, err := reqClaim.SignedString(JwyKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS

}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwyKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenHerder := context.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHerder == "" {
			//不存在
			code = errmsg.TokenError
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			context.Abort()
			return

		}
		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.TokenTypeError
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.TokenError
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.TokenTimeOut
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		context.Set("openid", key.Openid)
		context.Next()
	}
}