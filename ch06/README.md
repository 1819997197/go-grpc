## gRPC HTTP协议转换(一)

------

> * 定义了一个Hello Service，客户端发送包含字符串名字的请求，服务端返回Hello消息
> * 新增一个协议转换的gateway(gw.go)，把http协议请求转发给对应的grpc服务
> * 一个服务监听了两个端口: grpc服务监听50052端口；网关服务监听8080端口

#### 项目目录
```
$GOPATH/src/go-grpc/ch06

ch06
├── proto                                // proto描述文件
│   └── hello_http.proto
├── client.go                           // 客户端
├── server.go                           // 服务端
├── gw.go                               // 网关
└── README.md
// 目录结构有所调整，不想多次生成proto
```

#### Usage
1.Generate proto file
```
cd $GOPATH/src/go-grpc/ch06/proto

// 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

// 编译hello_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=go-grpc/ch06/proto/google/api:. ./*.proto

// 编译hello_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. ./*.proto
```


2.Run the service
```
cd $GOPATH/src/go-grpc/ch06/

// 编译grpc服务
go build -o server server.go
./server

// 编译gateway服务
go build -o gw gw.go
./gw
```

3.Run the client
```
cd $GOPATH/src/go-grpc/ch06
go run client.go

// curl/postman请求
curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
```