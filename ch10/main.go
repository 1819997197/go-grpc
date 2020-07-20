package main

import (
	"fmt"
	"go-grpc/ch10/frame"
	"go-grpc/ch10/register"
)

var name = "ws-order"
var port = ":8080"

func main() {
	// 1.load config
	if err := frame.LoadConfig("./config", "conf.yaml", "yaml"); err != nil {
		fmt.Println("load config err: ", err)
		return
	}

	// 2.run server
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
