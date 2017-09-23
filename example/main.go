package main

import (
	"github.com/clin003/tbsdk"
	"fmt"
)

func main(){
	p := tbsdk.OpenTbkTpwdParam{}
	p.Text = "hello tbsdk"
	p.Url = "https://www.taobao.com"
	p.User_id="2152088283"
	//fmt.Println(p.Text)
	//fmt.Println(p.APIName())
	//fmt.Println(p.ExtJSONParamName())
	//fmt.Println(p.ExtJSONParamValue())
	//tbsdk.UpdateKey("appkey","appsecret")
	//req ,nil:= tbsdk.Request(p)
	//fmt.Println(req)
	//fmt.Println(nil)
	req2,nil:=tbsdk.RequestWithKey("appkey","appsecret",p)
	fmt.Println(req2)
	fmt.Println(nil)
}
