package main

import (
	"fmt"
	pb "go-grpc/ch05/cli/proto"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/metadata"
)

const (
	// Address gRPC服务地址
	Address = "0.0.0.0:50052"
)

// customCredential 自定义认证(需要实现google.golang.org/grpc/credentials包下面的接口PerRPCCredentials)
type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return true
}

func deal(branch string) string {
	// TLS连接
	creds, err := credentials.NewClientTLSFromFile("./keys/crt.pem", "hello") //CN
	if err != nil {
		return "credentials.NewClientTLSFromFile err: " + err.Error()
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	// 指定自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	ctx := metadata.AppendToOutgoingContext(context.Background(), "trace-branch", branch)
	conn, err := grpc.DialContext(ctx, Address, opts...)
	if err != nil {
		return "grpc.DialContext err:" + err.Error()
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(ctx, req)
	if err != nil {
		return "c.SayHello err:" + err.Error()
	}

	return res.Message
}

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		branch := r.Header.Get("trace-branch")
		fmt.Println("branch: ", branch)
		msg := deal(branch)
		fmt.Println("deal msg: ", msg)
		w.Write([]byte(msg))
	})

	fmt.Println("listen 50050 before")
	http.ListenAndServe("0.0.0.0:50050", nil)
}
