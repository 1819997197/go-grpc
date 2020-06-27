## gRPC HTTP协议转换(二)

------

> * 定义了一个Hello Service，客户端发送包含字符串名字的请求，服务端返回Hello消息
> * 新增一个协议转换的gateway(gw.go)，把http协议请求转发给对应的grpc服务
> * grpc服务和gateway服务同时监听8080端口

#### 项目目录
```
$GOPATH/src/go-grpc/ch07

ch07
├── proto                                // proto描述文件
│   └── hello_http.proto
├── client.go                           // 客户端
├── server.go                           // grpc服务/gateway服务
└── README.md
```

#### Usage
1.Generate proto file
```
cd $GOPATH/src/go-grpc/ch07/proto

// 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

// 编译hello_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=go-grpc/ch07/proto/google/api:. ./*.proto

// 编译hello_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. ./*.proto
```


2.Run the service
```
cd $GOPATH/src/go-grpc/ch07/
go build -o server server.go
./server
```

3.Run the client
```
cd $GOPATH/src/go-grpc/ch07
go run client.go

// curl/postman请求
curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
```