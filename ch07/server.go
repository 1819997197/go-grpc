package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "go-grpc/ch07/proto"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

// 定义helloHTTPService并实现约定的接口
type helloHTTPService struct{}

// HelloHTTPService Hello HTTP服务
var HelloHTTPService = &helloHTTPService{}

// SayHello 实现Hello服务接口
func (h *helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

func main() {
	endpoint := "0.0.0.0:8080"
	// grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterHelloHTTPServer(grpcServer, HelloHTTPService)

	// gw server
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	gwmux := runtime.NewServeMux()
	if err := pb.RegisterHelloHTTPHandlerFromEndpoint(ctx, gwmux, endpoint, dopts); err != nil {
		fmt.Println("pb.RegisterHelloHTTPHandlerFromEndpoint err: ", err)
		return
	}

	// http服务
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	fmt.Println("gRPC and https listen on: ", endpoint)
	if err := http.ListenAndServe(":8080", grpcHandlerFunc(grpcServer, mux)); err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}

	return
}

// h2c或者cmux
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
	return handler
}
