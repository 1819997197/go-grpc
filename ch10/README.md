## gRPC 框架(http/grpc + config + 调用链 + gorm)

------

> * 基于grpc、gateway封装成一个可直接使用的grpc框架
> * 同时支持grpc、http协议，监听同一个端口
> * 支持多个proto文件
> * 自定义配置文件
> * APM

#### 项目目录
```
$GOPATH/src/go-grpc/ch10


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
