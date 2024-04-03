package main

import (
	"context"
	"go-learn/go_grpc/hello_grpc"
	pd "go-learn/go_grpc/hello_grpc"
	"go-learn/go_grpc/pubsub"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func main() {
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//注册服务
	hello_grpc.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	pd.RegisterPubsubServiceServer(grpcServer, NewPubsubService())
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("服务提供失败，err:%v", err)
		return
	}
}

// HelloServiceImpl 实现hello_grpc.go中的HelloServiceServer接口
type HelloServiceImpl struct {
	pd.UnimplementedHelloServiceServer
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pd.String) (*pd.String, error) {
	reply := &pd.String{Value: "hello" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream pd.HelloService_ChannelServer) error {
	//双向流特性
	for {
		//接收客户端的流数据
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		//对客户端的数据进行处理
		reply := &pd.String{Value: "hello:" + recv.GetValue()}

		//返回给客户端
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

// PubsubService 实现PubsubServiceServer接口
type PubsubService struct {
	pd.UnimplementedPubsubServiceServer
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Microsecond, 10),
	}
}

// Publish 发布消息到订阅者
func (p *PubsubService) Publish(ctx context.Context, arg *pd.String) (*pd.String, error) {
	p.pub.Publish(arg.GetValue())
	return &pd.String{}, nil
}

// Subscribe 通过Topic订阅
func (p *PubsubService) Subscribe(arg *pd.String, stream pd.PubsubService_SubscribeServer) error {
	//np := NewPubsubService()
	//fmt.Println(arg.GetValue())
	ch := p.pub.SubscriberTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			//fmt.Println("判断中", arg.GetValue())
			return strings.HasPrefix(key, arg.GetValue())
		}
		return false
	})
	for v := range ch {
		if err := stream.Send(&pd.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}
