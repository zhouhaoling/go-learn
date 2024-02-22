package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		go Consume(fmt.Sprintf("name%d", i))
	}

	go Publish()
	//阻塞主协程main
	fmt.Println("停止键入Ctrl+C")
	<-make(chan int)
}
