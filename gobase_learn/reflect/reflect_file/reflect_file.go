package main

import (
	"fmt"
)

// 解析文件和写入文件
func main() {
	//写入文件
	parseFile2("./gobase_learn/reflect/reflect_file/test.ini")
}

func parseFile(filename string) {

	var conf Config
	err := UnMarshalFile(filename, &conf)
	if err != nil {
		return
	}
	fmt.Printf("反序列化成功  conf: %#v\n  port: %#v\n", conf, conf.ServerConf.Port)
}

func parseFile2(filename string) {
	// 有一些假数据
	var conf Config
	conf.ServerConf.Ip = "127.0.0.1"
	conf.ServerConf.Port = 8000
	conf.MysqlConf.Port = 9000
	err := MarshalFile(filename, conf)
	if err != nil {
		return
	}
}
