package main

import (
	"context"
	pb "go-learn/go-grpc-tls/hello"
	"log"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

const address = "127.0.0.1:50052"

func main() {
	log.Println("客户端连接!")
	// TLS连接
	creds, err := credentials.NewClientTLSFromFile("./cert/server/server.pem", "test.example.com")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("err:", err)
	}
	defer conn.Close()
	// 初始化客户端
	c := pb.NewHelloServiceClient(conn)
	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Name)
}
