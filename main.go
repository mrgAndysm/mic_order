package main

import (
	"log"
	"prott/protoorder"
	"prott/service/grpc"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	//新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	srv := micro.NewService(
		//micro.Name服务端向consul注册服务名称 客户端需要通过该名称来进行调用
		micro.Name("go.micro.svr.order"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	srv.Init()
	protoorder.RegisterOrderServiceHandler(srv.Server(), new(grpc.OrderService))

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
