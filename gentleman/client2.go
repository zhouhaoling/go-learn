package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2/plugins/body"

	"gopkg.in/h2non/gentleman.v2"
)

type User struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	cli := gentleman.New()
	cli.URL("http://httpbin.org/post")
	//发送JSON格式数据
	//data := map[string]string{"foo": "bar"}
	//cli.Use(body.JSON(data))

	req := cli.Request()
	req.Method("POST")

	//发送xml格式数据
	u := User{Name: "dj", Age: 18}
	req.Use(body.XML(u))

	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}

	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}
