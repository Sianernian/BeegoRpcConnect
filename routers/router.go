package routers

import (
	"BeegoRpc/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{})
    beego.Router("/home", &controllers.MainController{})

}
