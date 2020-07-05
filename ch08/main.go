package main

import (
	"fmt"
	"go-grpc/ch08/frame"
	"go-grpc/ch08/register"
)

var name = "ws-order"
var port = ":8080"

func main() {
	application := &frame.GRPCApplication{
		App:                &frame.Application{Name: name},
		Port:               port,
		RegisterGRPCServer: register.RegisterGRPCServer,
		RegisterGateway:    register.RegisterGateway,
		RegisterHttpRoute:  register.RegisterHttpRoute,
	}
	if err := frame.Run(application); err != nil {
		fmt.Println("run err: ", err)
		return
	}
}
