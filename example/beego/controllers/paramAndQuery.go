package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) ParamsKey() {
	result := make(map[string]string)
	result["param id"] = c.Ctx.Input.Param(":id")
	result["user id"] = c.Ctx.Input.Param(":user_id")
	result["query id"] = c.Ctx.Input.Query("id")

	key := c.Ctx.Input.Params()
	fmt.Printf("%+v", key)
	c.Data["json"] = key
	c.ServeJSON()
}
