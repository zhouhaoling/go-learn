package main

import (
	"context"
	"fmt"
	pb "go-learn/go-grpc-tls/hello"
	"log"
	"net"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

const Address = "127.0.0.1:50052"

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (h *HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Name = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

func main() {
	fmt.Println("服务端启动")
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("./cert/server/server.pem", "./cert/server/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	if err != nil {
		log.Fatalf("failed to generate credentials %v", err)
	}
	//实例化grpc Server，并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))
	//注册HelloService
	pb.RegisterHelloServiceServer(s, &HelloService{})
	log.Println("Listen on " + Address + " with TLS")
	s.Serve(listen)
}
