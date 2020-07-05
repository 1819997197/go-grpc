package server

import (
	"context"
	pb "go-grpc/ch08/proto"
)

type helloHTTPService struct{}

var HelloHTTPService = &helloHTTPService{}

func (h *helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}
