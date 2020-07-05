## gRPC 框架(http/grpc + config)

------

> * 基于grpc、gateway封装成一个可直接使用的grpc框架
> * 同时支持grpc、http协议，监听同一个端口
> * 支持多个proto文件
> * 自定义配置文件

#### 项目目录
```
$GOPATH/src/go-grpc/ch09

ch09
├── proto                               // proto文件
│   └── hello_http.proto
├── frame                               // 框架封装
│   └── grpc.go
├── register                            // register http/grpc接口
│   └── register.go
├── server                              // 服务实现
│   └── hello_server.go
├── config                              // 配置文件
│   └── conf.yaml
├── main.go                             // 服务端入口文件
├── client.go                           // 客户端
└── README.md
```

#### Usage
1.Generate proto file
```
cd $GOPATH/src/go-grpc/ch09/proto

// 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

// 编译hello_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=go-grpc/ch09/proto/google/api:. ./*.proto

// 编译hello_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. ./*.proto
```


2.Run the service
```
cd $GOPATH/src/go-grpc/ch09/
go mod init go-grpc/ch09
go build -o srv main.go
./srv
```

3.Run the client
```
cd $GOPATH/src/go-grpc/ch09
go run client.go

// curl/postman请求
curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
```