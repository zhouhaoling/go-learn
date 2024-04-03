package main

import (
	"context"
	"fmt"
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

	//grpc简单使用
	//client := hello_grpc.NewHelloServiceClient(conn)
	//clientHello(client)
	//clientChannel(client)

	//grpc服务器流式使用
	client := pb.NewPubsubServiceClient(conn)
	clientPublish(client)
}

//获取信息

// clientPublish 服务器流式
func clientPublish(client pb.PubsubServiceClient) {
	ctx := context.Background()
	_, err := client.Publish(ctx, &pb.String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(ctx, &pb.String{Value: "docker: hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}

// grpc客户端简单使用
func clientHello(client pb.HelloServiceClient) {
	hello, err := client.Hello(context.Background(), &pb.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hello.GetValue())
}

// 双向数据流使用
func clientChannel(client pb.HelloServiceClient) {
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
