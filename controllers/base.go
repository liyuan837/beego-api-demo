package controllers

import (
	"github.com/astaxie/beego"
)

const RESCODE_OK  = 200
const RESCODE_FAIL = 201
const RESCODE_NOAUTH = 202

// HeroController operations for Hero
type BaseController struct {
	beego.Controller
}

//封装返回实体ResponseEntity
type ResponseEntity struct {
	Result  string      `json:"result" description:"返回说明"`
	Rescode int         `json:"rescode" description:"返回码"`
	Msg     string      `json:"msg" description:"描述"`
	Data    interface{} `json:"data" description:"数据"`
}

func(c *BaseController)  GetSuccessResult(data interface{},msg string) (ResponseEntity){
	return ResponseEntity{Rescode:RESCODE_OK, Result:"ok", Msg:msg, Data:data}
}

func(c *BaseController) GetFailureResult(msg string) (ResponseEntity) {
	return ResponseEntity{Rescode:RESCODE_FAIL, Result:"fail", Msg:msg}
}

func(c *BaseController) NoAuthResult(msg string) (ResponseEntity) {
	return ResponseEntity{Rescode:RESCODE_NOAUTH, Result:"no auth", Msg:msg}
}