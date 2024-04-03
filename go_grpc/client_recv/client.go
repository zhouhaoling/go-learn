package main

import (
	"context"
	"fmt"
	pb "go-learn/go_grpc/hello_grpc"
	"io"
	"log"

	"google.golang.org/grpc"
)

// 用来订阅信息
func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(context.Background(), &pb.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("接收完成")
				break
			}
			log.Fatal(err)
		}
		fmt.Println(recv.GetValue())
	}
}
