## gRPC 框架(http/grpc + config + 调用链 + gorm)

------

> * ws-order 对外提供rpc接口
> * ws-bff 裁剪、聚合后端服务，对外提供http接口

#### 项目目录
```
$GOPATH/src/go-grpc/ch12

ch12/
├── ws-order            // grpc服务
└── ws-bff              // http服务 裁剪、聚合后端rpc服务
```

#### Usage
1.Run the ws-order
```
详情见 $GOPATH/src/go-grpc/ch12/ws-order/README.md
```

2.Run the ws-bff
```
详情见 $GOPATH/src/go-grpc/ch12/ws-bff/README.md
```

3.Test
```
curl http://localhost:9100/
```