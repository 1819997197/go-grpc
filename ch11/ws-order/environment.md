# APM development environmentedit

## 整体架构
![整体架构](https://github.com/1819997197/go-grpc/blob/master/ch11/ws-order/apm.png)

## install-and-run
安装并运行总的有三种方式：

#### 1.Elastic Cloud上托管的Elasticsearch服务
```
https://www.elastic.co/cn/cloud/
```
#### 2.自己安装和管理
```
Step 1: Install Elasticsearch
Step 2: Install Kibana
Step 3: Install APM Server
Step 4: Install APM agents
Step 5: Configure APM
```

#### 3.使用docker快速搭建一个开发环境
```
https://www.elastic.co/guide/en/apm/get-started/7.8/quick-start-overview.html

安装中可能碰到的问题：
1.docker-compose.yml里面的镜像有的已经不存在，可以直接到hub.docker.com站点搜索对应的镜像并下载，并修改docker-compose.yml文件中对应的镜像名

2.服务启动不了
[root@will apm]# docker-compose up
Creating apm_elasticsearch_1 ... done

ERROR: for kibana  Container "2e3921db887c" is unhealthy.
ERROR: Encountered errors while bringing up the project.

查看容器具体错误：
[root@will ~]# docker logs 2e3921db887c
OpenJDK 64-Bit Server VM warning: UseAVX=2 is not supported on this CPU, setting it to UseAVX=1
OpenJDK 64-Bit Server VM warning: UseAVX=2 is not supported on this CPU, setting it to UseAVX=1

es资源不够导致，修改docker-compose.yml文件的es资源配置，尽量搞一个配置高一点的服务器
```

## 官方的安装文档
```
https://www.elastic.co/guide/en/apm/get-started/7.8/install-and-run.html
```

## 集成APM
集成APM有两种方式：

#### 1.使用官方集成包
```
// 一些web、数据库、日志、rpc框架已经内置
https://www.elastic.co/guide/en/apm/agent/go/1.x/builtin-modules.html#builtin-modules-apmhttp
```

#### 2.自定义集成
```
自定义集成涉及到三个概念：
a.Transaction 表示一次顶级会话(HTTP请求或RPC调用)
b.Span 在会话中的一次调用(数据库访问，或者对其他服务进行RPC调用))
c.Error 错误

https://www.elastic.co/guide/en/apm/agent/go/current/custom-instrumentation.html
```