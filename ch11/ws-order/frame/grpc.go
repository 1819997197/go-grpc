package frame

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"go.elastic.co/apm/module/apmgrpc"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net"
	"net/http"
	"strings"
)

type Application struct {
	Name string
}

type GRPCApplication struct {
	App                *Application
	Port               string
	GRPCServer         *grpc.Server
	GatewayServeMux    *runtime.ServeMux
	Mux                *http.ServeMux
	HttpServer         *http.Server
	ServerOptions      []grpc.ServerOption
	RegisterGRPCServer func(*grpc.Server) error
	RegisterGateway    func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error
	RegisterHttpRoute  func(*http.ServeMux) error
}

func initApplication(app *GRPCApplication) {
	app.GRPCServer = grpc.NewServer(grpc.UnaryInterceptor(apmgrpc.NewUnaryServerInterceptor()))
	app.GatewayServeMux = runtime.NewServeMux()

	mux := http.NewServeMux()
	mux.Handle("/", app.GatewayServeMux)
	app.Mux = mux

	app.HttpServer = &http.Server{
		Addr:    app.Port,
		Handler: grpcHandlerFunc(app.GRPCServer, app.Mux),
	}
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" { // CORS
				w.Header().Set("Access-Control-Allow-Origin", "*")
				headers := []string{"Content-Type", "Accept"}
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
				methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
				return
			}
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
	return handler
}

func Run(app *GRPCApplication) error {
	// 1.init application
	initApplication(app)

	// 2.register grpc and http
	if app.RegisterGRPCServer == nil {
		return fmt.Errorf("run app.RegisterGRPCServer is nil")
	}
	err := app.RegisterGRPCServer(app.GRPCServer)
	if err != nil {
		return fmt.Errorf("run app.RegisterGRPCServer err: %v", err)
	}

	if app.RegisterGateway != nil {
		err = app.RegisterGateway(context.Background(), app.GatewayServeMux, app.Port, []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			return fmt.Errorf("run app.RegisterGateway err: %v", err)
		}
	}

	if app.RegisterHttpRoute != nil {
		err = app.RegisterHttpRoute(app.Mux)
		if err != nil {
			return fmt.Errorf("run app.RegisterHttpRoute err: %v", err)
		}
	}

	// 3.start server
	conn, err := net.Listen("tcp", app.Port)
	if err != nil {
		return fmt.Errorf("TCP Listen err: %v", err)
	}
	fmt.Println("listen ", app.Port)
	err = app.HttpServer.Serve(conn)
	if err != nil {
		return fmt.Errorf("run serve err: %v", err)
	}

	return nil
}
