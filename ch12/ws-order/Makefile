
GOPATH:=$(shell go env GOPATH)
PROJECT:=go-grpc/ch12/ws-order
DIR:=${GOPATH}/src/${PROJECT}


.PHONY: proto
proto:
	protoc -I ${DIR}/proto --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:${DIR}/proto ${DIR}/proto/google/api/*.proto
	protoc -I ${DIR}/proto --go_out=plugins=grpc,Mgoogle/api/annotations.proto=${PROJECT}/proto/google/api:${DIR}/proto ${DIR}/proto/*.proto
	protoc -I ${DIR}/proto --grpc-gateway_out=logtostderr=true:${DIR}/proto ${DIR}/proto/*.proto

.PHONY: build
build: proto
	CGO_ENABLED=0 GOOS=linux go build -o service_order --ldflags "-extldflags -static" main.go

.PHONY: docker
docker:
	docker build -t service_order:0.1 .
