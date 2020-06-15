package jwtauth

import (
	"github.com/astaxie/beego"
	"github.com/jwt-go"
	"time"
)
//生成token
func GenerateToken(jwtUser JwtUser) (string,error ){
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["loginName"] = jwtUser.LoginName
	claims["userCode"] = jwtUser.UserCode
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(beego.AppConfig.String("jwtkey")))
	if err != nil {
		return "",err
	}
	return tokenString,nil
}

// 校验token是否有效
func CheckToken(token string) ( error, JwtUser) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwtkey")), nil
	})
	if err != nil {
		beego.Error("转换为jwt claims失败.", err)
		return err, JwtUser{}
	}
	claims,ok := t.Claims.(jwt.MapClaims)
	if !ok {
		beego.Error("get ParseToken claims error")
		return err, JwtUser{}
	}
	loginName := claims["loginName"].(string)
	userCode := claims["userCode"].(string)

	jwtUser := JwtUser{LoginName: loginName, UserCode: userCode}
	return nil, jwtUser
}
