package main

import (
	"fmt"
	"go-grpc/ch10/frame"
	"go-grpc/ch10/register"
	"go-grpc/ch10/vars"
)

var name = "ws-order"
var port = ":8080"

func main() {
	// 1.load config
	if err := frame.LoadConfig("./config", "conf.yaml", "yaml"); err != nil {
		fmt.Println("load config err: ", err)
		return
	}

	// 2.init mysql
	if err := initMysql(); err != nil {
		fmt.Println("init mysql err: ", err)
		return
	}

	// 3.run server
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

func initMysql() error {
	var err error
	vars.DB, err = frame.Instance()
	return err
}
