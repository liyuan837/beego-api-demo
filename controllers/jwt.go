package controllers

import (
	"beego-api-demo/jwtauth"
)

// HeroController operations for Hero
type JwtController struct {
	BaseController
}

// URLMapping ...
func (c *JwtController) URLMapping() {
	c.Mapping("login", c.Login)
}

// Post ...
// @Title Post
// @Description Login JwtUser
// @Success 200 {object} BaseController.ResponseEntity
// @Failure 201 body is empty
// @router /login [post]
func (c *JwtController) Login() {
	jwtUser := jwtauth.JwtUser{Id:1,LoginName:"liyuan",UserCode:"10001",Password:"123456"}
	token,err := jwtauth.GenerateToken(jwtUser)
	if err != nil {
		c.Data["json"] = c.NoAuthResult("登陆失败")
	}else{
		c.Data["json"] = c.GetSuccessResult(token,"登陆成功")
	}
	c.ServeJSON()
}

// Get ...
// @Title Get
// @Description get JwtInfo
// @Success 200 {object} BaseController.ResponseEntity
// @Failure 201 body is empty
// @router /get [post]
func (c *JwtController) GetJwtInfo() {
	token := c.Ctx.Request.Header.Get("Authorization")
	err,jwtUser := jwtauth.CheckToken(token)
	if err != nil {
		c.Data["json"] = c.NoAuthResult(err.Error())
	}else{
		c.Data["json"] = c.GetSuccessResult(jwtUser,"登陆成功")
	}
	c.ServeJSON()
}
