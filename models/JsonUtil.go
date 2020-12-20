package models

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type RpcJson struct {
	Id      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params []interface{} 	`json:"params"`
}

const RPCUSER  ="user" //rpc服务的用户名
const RPCPASSWORD  ="pwd" //rpc服务的用户名
const RPCPORT  ="http://127.0.0.1:8332" //rpc服务的用户名

func PraPareJson(method string,parmas ...interface{})string{
	Rpc :=&RpcJson{
		Id:      time.Now().UnixNano(),
		Jsonrpc: "5.2.0",
		Method:  method,
		Params:  parmas,
	}
	data,err:=json.Marshal(&Rpc)
	if err !=nil{
		log.Fatal(err.Error())
	}
	return string(data)
}

func ClienCount(Data string)*http.Response{
	client :=http.Client{};
	request, err :=http.NewRequest("POST",RPCPORT,strings.NewReader(Data))
	if err !=nil{
		log.Fatal(err.Error())
	}
	request.Header.Add("Encoding","UTF-8");
	request.Header.Add("Content-Type","application/json")
	request.SetBasicAuth(RPCUSER,RPCPASSWORD);
	resp,err :=client.Do(request);
	if err !=nil{
		log.Fatal(err.Error())
	}
	return resp
}

