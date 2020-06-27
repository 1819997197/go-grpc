package main

import (
	"fmt"
	"net"

	pb "go-grpc/ch06/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloHTTPService并实现约定的接口
type helloHTTPService struct{}

// HelloHTTPService 实现服务端接口
var HelloHTTPService = &helloHTTPService{}

// SayHello ...
func (h *helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloHTTPService
	pb.RegisterHelloHTTPServer(s, HelloHTTPService)

	fmt.Println("Listen on " + Address)
	if err := s.Serve(listen); err != nil {
		fmt.Println("err: ", err)
		return
	}
}
