package main

import (
	_ "beego-api-demo/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"beego-api-demo/controllers"
	"github.com/astaxie/beego/context"
	"beego-api-demo/components"
)
var JwtFilter = func(ctx *context.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token=="" && ctx.Request.RequestURI != "/v1/jwt/login" {
		base := controllers.BaseController{}
		ctx.Output.JSON(base.NoAuthResult("请先登录"), false, false)
	}
}
func init() {
	//配置数据源、连接池
	components.InitMysql()

	//配置redis连接池
	components.InitRedis()

	//注册日志,调用logger初始化
	components.InitLogger()

	//注册登录拦截器
	beego.InsertFilter("/*",beego.BeforeRouter, JwtFilter)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}

