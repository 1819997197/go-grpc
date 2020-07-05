## gRPC 框架

------

> * 基于grpc、gateway封装成一个可直接使用的grpc框架(对前面几节的知识点做一个封装)
> * 同时支持grpc、http协议，监听同一个端口
> * 支持多个proto文件

#### 项目目录
```
$GOPATH/src/go-grpc/ch08

ch08
├── proto                               // proto文件
│   └── hello_http.proto
├── frame                               // 框架封装
│   └── grpc.go
├── register                            // register http/grpc接口
│   └── register.go
├── server                              // 服务实现
│   └── hello_server.go
├── main.go                             // 服务端入口文件
├── client.go                           // 客户端
└── README.md
```

#### Usage
1.Generate proto file
```
cd $GOPATH/src/go-grpc/ch08/proto

// 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

// 编译hello_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=go-grpc/ch08/proto/google/api:. ./*.proto

// 编译hello_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. ./*.proto
```


2.Run the service
```
cd $GOPATH/src/go-grpc/ch08/
go mod init go-grpc/ch08
go build -o srv main.go
./srv
```

3.Run the client
```
cd $GOPATH/src/go-grpc/ch08
go run client.go

// curl/postman请求
curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
```