package components

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitMysql() (err error){
	mysqlInfo := beego.AppConfig.String("mysqlinfo")
	orm.RegisterDataBase("default", "mysql", mysqlInfo)
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)
	beego.Info("mysql初始化成功")
	return
}
