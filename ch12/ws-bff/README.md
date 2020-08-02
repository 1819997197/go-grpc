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
#### Usage
1.Generate the service
```
// 生成proto并编译出二进制文件
make build
```

2.构建镜像
```
make docker

// 测试镜像是否可用
[root@will ws-order]# docker run --name ws_bff1 -d -p 9100:9100 service_bff:0.1
4691cbef1824c20c7863df0995fffbae24ba677225cafa2873cc423ea77e71f4
[root@will ws-order]# curl localhost:9100
it ok!Hello gRPC.
```

3.在k8s中把服务跑起来
```
kubectl apply -f bff-deployment.yaml -f bff-svc.yaml
```