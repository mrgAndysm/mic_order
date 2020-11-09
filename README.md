# 基于go-micro 实现一个订单微服务
    github.com/go-micro/micro/v2

### 注意： 
    基于v2版本，注意v3版本默认不支持consul
# protoc 安装

依赖：autoconf automake libtool curl make g++ unzip libffi-dev

    下载：https://github.com/protocolbuffers/protobuf

    cd protoc

    ./autogen.sh

    ./configure

    make && make install
    
    ldconfig //刷新共享库
    
    测试：protoc -h 

# gen-go-pro安装:
    go get -v -u github.com/golang/protobuf/proto

    cd $GOPATH/src/github/golang/protobuf/gen-go-pro/

    go bilud

    cp gen-go... /usr/bin

# protoc-gen-micro安装：
    go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master

    protoc --proto_path=$(shell pwd)/proto/ --micro_out=$(shell pwd)/protoorder --go_out=$(shell pwd)/protoorder $(shell pwd)/proto/*proto

# docker - consul部署

leader节点

    docker run -d -p 8500:8500 -v /data/consul:/consul/data -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_1 4f7b214361a7 agent -server -bootstrap -ui -node=1 -client='0.0.0.0'

集群从节点

    docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_2 4f7b214361a7 agent -client -node=clint -join='172.17.0.2' -client='0.0.0.0'

    docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_3 4f7b214361a7 agent -client -node=clint2 -join='172.17.0.2' -client='0.0.0.0'

consul服务界面

    http://192.168.33.12:8500/
# 安装成功后
 执行：    make run

### 偶遇问题：
    etcd undefined: resolver.BuildOption
    需要将grpc1.27.0 替换成grpc1.26.0版本
    google.golang.org/grpc v1.27.0
    替换成
    google.golang.org/grpc v1.26.0
    
    具体步骤:
    
    先删除pkg/mod: rm -rf  或手动删除
    再替版本: go mod edit -require=google.golang.org/grpc@v1.26.0
    下载指定版本v1.26.0: go get -u -x google.golang.org/grpc@v1.26.0
    然后再go mod init
    再运行程:make run
    