package main

import (
	"context"
	"fmt"
	"go-learn/go_grpc/hello_grpc"
	pb "go-learn/go_grpc/hello_grpc"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := hello_grpc.NewHelloServiceClient(conn)
	hello, err := client.Hello(context.Background(), &pb.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hello.GetValue())

	//接收和发送数据流
	stream, err := client.Channel(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err = stream.Send(&pb.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(recv.GetValue())
	}
}
