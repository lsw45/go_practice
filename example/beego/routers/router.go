package routers

import "github.com/astaxie/beego"
import "../controllers"

func init() {
	beego.Router("/:id/:user_id", &controllers.MainController{}, "*:ParamsKey")
}
