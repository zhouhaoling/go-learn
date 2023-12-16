package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
)

func main() {
	//创建http客户端
	cli := gentleman.New()
	//设置要请求的地址
	cli.URL("https://dog.ceo")
	//创建一个请求对象
	req := cli.Request()
	//设置请求的路径
	req.Path("/api/breeds/image/random")
	//设置请求的头部
	req.SetHeader("Client", "gentleman")
	//发送请求
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request err:%v\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %dn", res.StatusCode)
		return
	}

	fmt.Printf("Body: %s", res.String())
}
