package main

import (
	"fmt"
	"net"

	pb "go-grpc/ch04/server/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/metadata"
)

const (
	// Address gRPC服务地址
	Address = "0.0.0.0:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = &helloService{}

// SayHello 实现Hello服务接口
func (h *helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx) //获取grpc协议头里面的信息
	if ok {
		for key, value := range md {
			for k, v := range value {
				fmt.Println("md", key, k, v)
			}
		}
	}

	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("v1: Hello %s", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("./keys/crt.pem", "./keys/privatekey.pem")
	if err != nil {
		fmt.Println("Failed to generate credentials ", err)
		return
	}

	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	fmt.Println("Listen on " + Address + " with TLS")
	if err := s.Serve(listen); err != nil {
		fmt.Println("run server err: ", err)
		return
	}
}
