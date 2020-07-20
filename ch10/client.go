package main

import (
	"fmt"
	pb "go-grpc/ch10/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "0.0.0.0:8080"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial err: ", err)
		return
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloHTTPClient(conn)

	// 调用方法
	reqBody := new(pb.HelloHTTPRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		fmt.Println("c.SayHello err: ", err)
		return
	}

	fmt.Println("message", r.Message)
}

// OR: curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
