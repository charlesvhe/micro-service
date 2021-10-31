package main

import (
	"context"
	"fmt"
	"log"
	"time"

	gRpcClient "github.com/asim/go-micro/plugins/client/grpc/v3"
	gRpcServer "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/v3"

	"github.com/asim/go-micro/plugins/registry/nacos/v3"
	"github.com/asim/go-micro/v3/registry"

	"provider/proto"
)

type Provider struct{}

func (p *Provider) Test(ctx context.Context, req *proto.Msg, rsp *proto.Msg) error {
	rsp.Name = "Hello " + req.Name + "! " + time.Now().String()
	fmt.Println(rsp.Name)

	return nil
}

func main() {
	registry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"192.168.1.9:8848"}
		options.Context = context.Background()
	})
	// create a new service
	service := micro.NewService(
		micro.Server(gRpcServer.NewServer()),
		micro.Client(gRpcClient.NewClient()),
		micro.Name("go-provider"),
		micro.Address(":5000"),
		micro.Registry(registry),
	)

	// initialise flags
	service.Init()
	// gRPC server
	proto.RegisterProviderHandler(service.Server(), new(Provider))

	// start the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
