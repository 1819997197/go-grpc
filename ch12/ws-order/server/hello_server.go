package server

import (
	"context"
	pb "go-grpc/ch12/ws-order/proto"
)

type helloHTTPService struct{}

var HelloHTTPService = &helloHTTPService{}

func (h *helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
	//list, err := repository.NewUserModel().FindList(ctx)
	//if err != nil {
	//	fmt.Println("SayHello err: ", err)
	//	return resp, nil
	//}
	//if list == nil {
	//	fmt.Println("list is nil")
	//	return resp, nil
	//}
	//
	//for _, v := range list {
	//	fmt.Println("user: ", v.Id, "---", v.Name)
	//}
	//
	//return resp, nil
}
