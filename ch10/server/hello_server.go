package server

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	pb "go-grpc/ch10/proto"
)

type helloHTTPService struct{}

var HelloHTTPService = &helloHTTPService{}

func (h *helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	fmt.Println(dsn)

	return resp, nil
}
