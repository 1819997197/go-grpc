## go-grpc

#### 一、安装ProtocolBuffers 3.0.0-beta-3或更高版本
```
1.下载链接: https://github.com/protocolbuffers/protobuf/releases
2.找到对应的版本下载并解压
3.安装
./configure
make
make install
4.查看是否看着正常
[vagrant@localhost ch01]$ protoc --version
libprotoc 3.6.1
```

#### 二、安装 ProtoBuf 相关的 golang 依赖库
```
1.开启go mod模式: export GO111MODULE=on GOPROXY=https://goproxy.cn
2.写go文件
package ch01

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
)

3.增加确实的包: go mod tidy

4.go install \
      github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
      github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
      github.com/golang/protobuf/protoc-gen-go

5.会在$GOBIN目录生成三个二进制文件(protoc-gen-go、protoc-gen-swagger、protoc-gen-grpc-gateway)
```

#### 三、编译器的使用
使用protoc命令编译.proto文件,不同语言支持需要指定输出参数，如:
```
protoc --proto_path=IMPORT_PATH --cpp_out=DST_DIR --java_out=DST_DIR --python_out=DST_DIR
--go_out=DST_DIR --ruby_out=DST_DIR --javanano_out=DST_DIR --objc_out=DST_DIR
--csharp_out=DST_DIR path/to/file.proto

-I 参数：指定import路径，可以指定多个-I参数，编译时按顺序查找，不指定时默认查找当前目录

--go_out：golang编译支持，支持以下参数
        plugins=plugin1+plugin2 - 指定插件，目前只支持grpc，即：plugins=grpc
        M 参数 - 指定导入的.proto文件路径编译后对应的golang包名(不指定本参数默认就是.proto文件中import语句的路径)
        import_prefix=xxx - 为所有import路径添加前缀，主要用于编译子目录内的多个proto文件，这个参数按理说很有用，尤其适用替代一些情况时的M参数，但是实际使用时有个蛋疼的问题导致并不能达到我们预想的效果，自己尝试看看吧
        import_path=foo/bar - 用于指定未声明package或go_package的文件的包名，最右面的斜线前的字符会被忽略
        末尾 :编译文件路径 .proto文件路径(支持通配符)
```