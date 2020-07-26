## ws-order(http/grpc + config + 调用链 + gorm)

------

> * 基于grpc、gateway封装成一个可直接使用的grpc框架
> * 同时支持grpc、http协议，监听同一个端口
> * 支持多个proto文件
> * 自定义配置文件
> * APM

#### 项目目录
```
$GOPATH/src/go-grpc/ch11/ws-order

ch11/ws-order
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

#### 安装APM
```
//本示例中采用了environment.md文件中第三种安装方式(部分镜像有所更改)
1.把docker-compose.yml拷贝到安装了docker、docker-compose的服务器上面

2.docker-compose up
// ctrl+c will stop all of the containers

3.成功启动之后会显示如下信息
[root@will apm]# docker-compose up
Starting apm_elasticsearch_1 ... done
Starting apm_kibana_1        ... done
Starting apm_apm-server_1    ... done

4.运行成功之后，根据引导配置项目(http://localhost:5601/app/kibana#/home/tutorial/apm)
export ELASTIC_APM_SERVER_URL=http://192.168.1.106:8200
export ELASTIC_APM_SERVICE_NAME=ws

5.通过 http://localhost:5601/app/apm 查看调用链数据
```

#### Usage
1.创建数据库表，修改配置文件
```
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

// 配置文件 /config/conf.yaml
```

2.配置APM环境变量
```
export ELASTIC_APM_SERVER_URL=http://192.168.1.106:8200
export ELASTIC_APM_SERVICE_NAME=ws-order
```

3.Generate proto file
```
cd $GOPATH/src/go-grpc/ch11/ws-order/proto

// 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

// 编译hello_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=go-grpc/ch11/ws-order/proto/google/api:. ./*.proto

// 编译hello_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. ./*.proto
```

4.Run the service
```
cd $GOPATH/src/go-grpc/ch11/ws-order/
go mod init go-grpc/ch11/ws-order
go build -o srv main.go
./srv
```

5.Test the ws-order
```
cd $GOPATH/src/go-grpc/ch11/ws-order
go run client.go

// curl/postman请求
curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
```