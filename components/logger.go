package components

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)
const (
	LevelEmergency = iota  //0
	LevelAlert        //1
	LevelCritical	  //2
	LevelError		  //3
	LevelWarning	  //4
	LevelNotice		  //5
	LevelInformational//6
	LevelDebug		  //7
)

func InitLogger() (err error){
	BConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil{
		fmt.Println("config init error:", err)
		return
	}
	maxlines, lerr := BConfig.Int64("log::maxlines")
	if lerr != nil {
		maxlines = 10000
	}

	logConf := make(map[string]interface{})
	logConf["filename"] = BConfig.String("log::log_path")
	level,_ := BConfig.Int("log::log_level")
	logConf["level"] = level
	logConf["maxlines"] = maxlines

	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	beego.SetLogger(logs.AdapterFile, string(confStr))
	beego.SetLogFuncCall(true)
	return
}