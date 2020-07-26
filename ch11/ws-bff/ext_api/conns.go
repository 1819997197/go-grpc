package ext_api

import (
	"context"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
)

const (
	Ws_Order_Address = "0.0.0.0:8080"
)

func GetWsOrderClient(ctx context.Context) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, Ws_Order_Address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
