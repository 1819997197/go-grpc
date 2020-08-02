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
[root@will ws-bff]# kubectl get pods
NAME                                READY   STATUS    RESTARTS   AGE
bff-deployment-564948b66f-9kd8f     1/1     Running   0          3m18s
bff-deployment-54f57cf6bf-gfblf   	1/1     Running   0          3m18s
order-deployment-54f57cf6bf-st6cc   1/1     Running   0          26m
order-deployment-55f8b9f87b-th4cb   1/1     Running   0          26m
[root@will ws-bff]# kubectl get svc -o wide
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE     SELECTOR
bff-svc      ClusterIP   10.107.56.215   <none>        9100/TCP   2m38s   app=bff-deployment
order-svc    ClusterIP   10.104.88.142   <none>        8080/TCP   24m     app=order-deployment
[root@will ws-bff]# curl 10.107.56.215:9100/
it ok!Hello gRPC.
```