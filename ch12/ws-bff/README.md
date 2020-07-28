## ws-bff(gin + grpc +调用链)

------

> * 裁剪、聚合后端服务
> * 对外提供http接口

#### 项目目录
```
$GOPATH/src/go-grpc/ch12/ws-bff

ch12/ws-bff
├── ext_api                         // grpc client
│   ├── conns.go
│   └── ws_order_client.go
├── handler                         // 业务层方法
│   └── handler.go
├── proto                           // proto文件
├── main.go                         // 服务入口文件
└── README.md
```

#### Usage
1.配置APM环境变量
```
export ELASTIC_APM_SERVER_URL=http://192.168.1.106:8200
export ELASTIC_APM_SERVICE_NAME=ws-bff
```

3.Generate proto file
```
cd $GOPATH/src/go-grpc/ch12/ws-order/proto

// 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

// 编译hello_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=go-grpc/ch12/ws-bff/proto/google/api:. ./*.proto

// 编译hello_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. ./*.proto
```

4.Run the service
```
cd $GOPATH/src/go-grpc/ch12/ws-bff/
go mod init go-grpc/ch12/ws-bff
go build -o srv main.go
./srv
```

5.Test the ws-order
```
// curl/postman请求
curl http://localhost:9100/
```