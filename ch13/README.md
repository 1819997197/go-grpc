# envoy

## 生成envoy配置文件
```
1.设置环境变量
export envoy_proxy_port="10004" //envoy监听端口
export envoy_branch="f1,master" //分支(多个分支用逗号隔开)
export envoy_service_name="order_bff" //bff/服务名
export envoy_service_port="8080" //服务端口
export envoy_ssl="1" //协议是否加密(可选, 1:加密 其它:非加密 默认不加密)
export envoy_protocol="grpc" //协议(http/grpc)
export envoy_certificate_chain="/etc/conf/crt.pem" //证书目录(envoy_ssl=1时，必填)
export envoy_private_key="/etc/private_key.pem" //私钥目录(envoy_ssl=1时，必填)

2.利用generate_config.go文件生成envoy.yaml配置文件
```

## docker构建&运行测试
```
1.构建
docker build -t envoy-proxy:0.1 .

2.测试
docker run -e envoy_proxy_port=10000 -e envoy_branch="f1,master" -e envoy_service_name="order_bff" -e envoy_service_port=8080 -e envoy_protocol="http" --rm -it envoy-proxy:0.1
```