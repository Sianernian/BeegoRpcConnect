package controllers

import (
	"BeegoRpc/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
)

type MainController struct {
	beego.Controller
}

var ch chan interface{}
// 刷新页面
func (c *MainController)Get(){
	c.TplName = "home.html"
}

func (c *MainController) Post() {
	JsonData := models.PraPareJson("getblockhash",1)
	DataString := models.ClienCount(JsonData)

	Code := DataString.StatusCode;
	if Code == 200 {
		bytes, err := ioutil.ReadAll(DataString.Body)
		if err != nil {
			c.Ctx.WriteString("解析错误")
		}
		var f models.RPCResult
		err = json.Unmarshal(bytes, &f);
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(f.Result)
		fmt.Println(f)
		c.Data["result"]= f
	//	c.ServeJSON()
		c.TplName = "home.html"
	} else {
		fmt.Println("请求失败")
	}

}

//func (c *MainController)Post() {
//	//解析数据
//
//}
