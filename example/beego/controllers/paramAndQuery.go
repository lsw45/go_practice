package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) ParamsKey() {
	result := make(map[string]string)
	result["param id"] = c.Ctx.Input.Param(":id")
	result["user id"] = c.Ctx.Input.Param(":user_id")
	result["query id"] = c.Ctx.Input.Query("id")
	c.Data["json"] = result
	c.ServeJSON()
}
