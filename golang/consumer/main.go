package main

import (
	"context"
	"fmt"
	"log"

	gRpcClient "github.com/asim/go-micro/plugins/client/grpc/v3"
	gRpcServer "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/v3"

	"github.com/asim/go-micro/plugins/registry/nacos/v3"
	"github.com/asim/go-micro/v3/registry"

	"consumer/proto"
)

type Consumer struct{}

var _providerClient proto.ProviderService

func (p *Consumer) Test(ctx context.Context, req *proto.Msg, rsp *proto.Msg) error {
	res, err := _providerClient.Test(context.Background(), &proto.Msg{
		Name: req.Name,
	})

	if err != nil {
		return err
	}

	rsp.Name = "Hello consumer " + res.Name + "! "
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
		micro.Name("go-consumer"),
		micro.Address(":60000"),
		micro.Registry(registry),
	)

	// initialise flags
	service.Init()
	// gRPC server
	proto.RegisterConsumerHandler(service.Server(), new(Consumer))
	// gRPC client
	_providerClient = proto.NewProviderService("go-provider", service.Client())

	// start the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
