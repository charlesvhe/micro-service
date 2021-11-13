package main

import (
	"bff/proto"
	"context"
	"log"

	gRpcClient "github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/isfk/go-micro-plugins/registry/nacos/v3"

	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
)

var _providerClient proto.ProviderService
var _consumerClient proto.ConsumerService

func main() {

	srv := http.NewServer(
		server.Name("go-bff"),
		server.Address(":8888"),
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// register router
	demo := newDemo()
	demo.InitRouter(router)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	registry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"localhost:8848"}
		options.Context = context.Background()
	})
	service := micro.NewService(
		micro.Server(srv),
		micro.Client(gRpcClient.NewClient()),
		micro.Registry(registry),
	)
	// initialise flags
	service.Init()

	// gRPC client
	_providerClient = proto.NewProviderService("go-provider", service.Client())
	_consumerClient = proto.NewConsumerService("go-consumer", service.Client())

	// start the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type Test struct{}

func newDemo() *Test {
	return &Test{}
}

func (a *Test) InitRouter(router *gin.Engine) {
	router.GET("/provider", a.provider)
	router.GET("/consumer", a.consumer)
}

func (a *Test) provider(c *gin.Context) {
	res, err := _providerClient.Test(context.Background(), &proto.Msg{
		Name: "CHE",
	})

	if err != nil {
		c.JSON(500, gin.H{"msg": "call provider err"})
	}

	c.JSON(200, gin.H{"name": res.Name})
}

func (a *Test) consumer(c *gin.Context) {
	res, err := _consumerClient.Test(context.Background(), &proto.Msg{
		Name: "CHE",
	})

	if err != nil {
		c.JSON(500, gin.H{"msg": "call consumer err"})
	}

	c.JSON(200, gin.H{"name": res.Name})
}
