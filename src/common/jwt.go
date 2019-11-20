package common

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

const SecretKey  ="secretkey"


type jwtCustomClaims struct {
	jwt.StandardClaims
	Uername string `json:"username"`
	Admin bool `json:"admin"`
}

// 刷新token
func NewCookie(username string, isAdmin bool)(newCookie string, err error) {
	newToken, err := CreateToken([]byte(SecretKey), "sgst", username, isAdmin)
    if err != nil {
    	logs.Error(err)
    	return "", err
	}
	cookie := http.Cookie {
		Name: "token",
		Value: newToken,
		HttpOnly: true,
	}
	return cookie.String(), nil
}
// 构建token字符串
func CreateToken(SecretKey []byte, issuer string, username string, isAdmin bool) (tokenString string, err error) {				//创建token
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,
		},
		username,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)

	return
}
// 解析token
func ParseToken(tokenSrt string, SecretKey []byte) (username string, err error) {	//解析token
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims := token.Claims
	logs.Info("claims:", claims)
	logs.Info(fmt.Sprintf("claims uid:%v,claims uid type:%T\n", claims.(jwt.MapClaims)["username"],claims.(jwt.MapClaims)["username"]))

	username = claims.(jwt.MapClaims)["username"].(string)
	return username, nil
}
//func TestCreateToken(t *testing.T)  {					//测试
//	token, _ := CreateToken([]byte(SecretKey), "YDQ", 2222, true)
//
//	fmt.Println(token)
//	claims, err := ParseToken(token, []byte(SecretKey))
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//
//	fmt.Println("claims:", claims)
//	fmt.Printf("claims uid:%v,claims uid type:%T\n", claims.(jwt.MapClaims)["uid"],claims.(jwt.MapClaims)["uid"])
//
//	//此处注意接口(interface)的取值方法
//	var i float64
//
//	//错误
//	//i=claims.(jwt.MapClaims)["uid"]
//
//	//正确
//	i=claims.(jwt.MapClaims)["uid"].(float64)
//	fmt.Printf("i:%v,i type:%T\n",i,i)
//}