package register

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "go-grpc/ch10/proto"
	"go-grpc/ch10/server"
	"google.golang.org/grpc"
	"net/http"
)

// 注册pb的Server
func RegisterGRPCServer(grpcServer *grpc.Server) error {
	pb.RegisterHelloHTTPServer(grpcServer, server.HelloHTTPService)

	return nil
}

// 注册pb的Gateway
func RegisterGateway(ctx context.Context, gateway *runtime.ServeMux, endPoint string, dopts []grpc.DialOption) error {
	if err := pb.RegisterHelloHTTPHandlerFromEndpoint(ctx, gateway, endPoint, dopts); err != nil {
		return err
	}

	return nil
}

// 注册http接口 func funcName(w http.ResponseWriter, r *http.Request){}
func RegisterHttpRoute(serverMux *http.ServeMux) error {
	return nil
}
