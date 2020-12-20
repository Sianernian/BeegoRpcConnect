package main

import (
	_ "BeegoRpc/routers"
	"github.com/astaxie/beego"
)

func main() {
	//db_mysql.ConnectDB();

	beego.SetStaticPath("/views","./views")
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()
}
