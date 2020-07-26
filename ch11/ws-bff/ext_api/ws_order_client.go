package ext_api

import (
	"fmt"
	pb "go-grpc/ch11/ws-bff/proto"
	"golang.org/x/net/context"
)

func SayHello(ctx context.Context) (string, error) {
	conn, err := GetWsOrderClient(ctx)
	if err != nil {
		fmt.Println("SayHello GetWsOrderClient err:", err) //需要记录日志
		return "", err
	}
	defer conn.Close()

	cli := pb.NewHelloHTTPClient(conn) // 初始化客户端
	reqBody := new(pb.HelloHTTPRequest)
	reqBody.Name = "gRPC"
	r, err := cli.SayHello(ctx, reqBody) // 调用方法
	if err != nil {
		fmt.Println("SayHello cli.SayHello err: ", err)
		return "", err
	}

	return r.Message, nil
}
