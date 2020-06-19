## Hello gRPC + TLS

------

> * 定义了一个Hello Service，客户端发送包含字符串名字的请求，服务端返回Hello消息
> * 从http头信息里面获取trace-branch字段信息，并在grpc协议头里面设置trace-branch并传递到服务端
> * 基于SSL/TLS认证方式，证书创建方式见ch03

#### 项目目录
```
$GOPATH/src/go-grpc/ch04

ch04
├── cli                     // 客户端
│   ├── keys                // 证书文件
│   ├── main.go
│   └── proto               // proto描述文件
│       ├── hello.pb.go
│       └── hello.proto
├── server                  // 服务端
│   ├── keys                // 证书文件
│   ├── main.go
│   └── proto               // proto描述文件
│       ├── hello.pb.go
│       └── hello.proto
└── README.md
```

#### Usage

1.Run the service
```
cd $GOPATH/src/go-grpc/ch04/server/
protoc -I . --go_out=plugins=grpc:. ./proto/*.proto
go mod init go-grpc/ch04/server
go mod tidy
go build -o server
./server
```

2.Run the client
```
cd $GOPATH/src/go-grpc/ch04/cli/
protoc -I . --go_out=plugins=grpc:. ./proto/*.proto
go mod init go-grpc/ch04/cli
go mod tidy
go build -o cli
./cli
```

3.Test
```
// curl/postman请求
curl -H "trace-branch:master" localhost:50050/test

// 客户端打印
[vagrant@localhost cli]$ ./cli
listen 50050 befor
branch:  master
deal msg:  v1: Hello gRPC

// 服务端打印
[vagrant@localhost server]$ ./server
Listen on 0.0.0.0:50052 with TLS
md trace-branch 0 master
md :authority 0 hello
md content-type 0 application/grpc
md user-agent 0 grpc-go/1.29.1

```