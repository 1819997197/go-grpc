## ws-order(http/grpc + config + 调用链 + gorm)

------

> * 基于grpc、gateway封装成一个可直接使用的grpc框架
> * 同时支持grpc、http协议，监听同一个端口
> * 支持多个proto文件
> * 自定义配置文件
> * APM

#### 项目目录
```
$GOPATH/src/go-grpc/ch12/ws-order

ch12/ws-order
├── frame                           // 框架封装
│   ├── config.go                   // init config
│   ├── grpc.go                     // grpc/http
│   └── mysql.go                    // init mysql
├── register                        // register http/grpc接口
│   └── register.go
├── repository                      // 仓储层(封装操作数据库方法)
│   └── user.go
├── server                          // 服务实现
│   └── hello_server.go
├── vars                            // 全局变量
│   └── varsiable.go
├── config                          // 配置文件
│   └── conf.yaml
├── proto                           // proto文件
├── docker-compose.yml              // APM docker-compose文件
├── environment.md                  // APM安装方法介绍
├── main.go                         // 服务端入口文件
├── client.go                       // 客户端
└── README.md
```

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
[root@will ws-order]# docker run --name ws2 -d -p 8080:8080 service_order:0.1
1e94eec4089c18bb847c52ab0a7a2c9f21da7015f0d410396fd9232a64e97a23
[root@will ws-order]# docker logs 1e94
listen  :8080
[root@will ws-order]# curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
{"message":"Hello gRPC-HTTP is working!."}
```

3.在k8s中把服务跑起来
```
kubectl apply -f order-deployment.yaml -f order-svc.yaml
```