package main

import (
	"fmt"
	pb "go-grpc/ch06/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// 连接
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial err", err)
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

	fmt.Println(r.Message)
}

// OR: curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
