package main

import (
	"fmt"
	"net"

	pb "go-grpc/ch05/server/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

func auth(ctx context.Context) error {
	fmt.Println("auth interceptor")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("auth 无Token认证信息")
		return fmt.Errorf("invalid code: %v", codes.Unauthenticated)
	}

	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "101010" || appkey != "i am key" {
		fmt.Printf("Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
		return fmt.Errorf("invalid code: %v", codes.Unauthenticated)
	}

	return nil
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

	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(creds))

	// 注册interceptor
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = auth(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	// 实例化grpc Server
	s := grpc.NewServer(opts...)

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	fmt.Println("Listen on " + Address + " with TLS")
	if err := s.Serve(listen); err != nil {
		fmt.Println("run server err: ", err)
		return
	}
}
