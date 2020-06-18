## Hello gRPC

------

hello demo定义了一个Hello Service，客户端发送包含字符串名字的请求，服务端返回Hello消息。具体流程如下：

> * 编写.proto描述文件
> * 编译生成.pb.go文件
> * 编写服务端，实现proto文件定义的接口
> * 编写客户端，调用服务端接口

#### 项目目录
```
$GOPATH/src/go-grpc/ch02

ch02
├── cli                     // 客户端
│   ├── main.go
│   └── proto
│       ├── hello.pb.go
│       └── hello.proto
├── server                  // 服务端
│   ├── main.go
│   └── proto
│       ├── hello.pb.go
│       └── hello.proto
└──  README.md
```

#### Usage

1.Run the service
```
cd $GOPATH/src/go-grpc/ch02/server/
protoc -I . --go_out=plugins=grpc:. ./proto/*.proto
go build -o server
./server
```

2.Run the client
```
cd $GOPATH/src/go-grpc/ch02/cli/
protoc -I . --go_out=plugins=grpc:. ./proto/*.proto
go build -o cli
./cli
```