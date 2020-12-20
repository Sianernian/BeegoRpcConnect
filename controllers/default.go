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


// 刷新页面
func (c *MainController) Get() {
	c.TplName = "home.html"
}

var ch chan interface{}


func (c *MainController)Post() {
	//解析数据
	JsonData := models.PraPareJson("getbestblockhash")
	DataString := models.ClienCount(JsonData)
	Code := DataString.StatusCode;
	if Code == 200 {
		bytes, err := ioutil.ReadAll(DataString.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		var f models.RPCResult
		err = json.Unmarshal(bytes, &f);
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(f.Result)
		c.Data["Result"]= f.Result
		c.TplName = "home.html"
	} else {
		fmt.Println("请求失败")
	}
}
