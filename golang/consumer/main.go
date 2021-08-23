package main

import (
	"context"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	"github.com/isfk/go-micro-plugins/registry/nacos/v3"
	"github.com/micro/micro/v3/service/logger"
	"micro-service/golang/consumer/handler"
	"os"
)

var etcdRegistry registry.Registry

const nacosNamespace = "dev"

const (
	defaultNacosAddr      = "127.0.0.1:8848"
	defaultNacosNamespace = "dev"
)

func main() {
	// 从环境变量中获取nacos的ip和port
	var nacosAddr string
	nacosAddr = os.Getenv("NacosAddr")
	if nacosAddr == "" {
		nacosAddr = defaultNacosAddr
	}

	var nacosNamespace string
	nacosNamespace = os.Getenv("NacosNamespace")
	if nacosNamespace == "" {
		nacosNamespace = defaultNacosNamespace
	}
	nacosRegistry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{nacosAddr}
		// 支持 namespace
		options.Context = context.WithValue(context.Background(), &nacos.NacosNamespaceContextKey{}, nacosNamespace)

	})
	srv := httpServer.NewServer(
		server.Name("consumer"),
		server.Address(":8080"),
	)

	router := gin.Default()
	// 注册router
	sumHandler := handler.NewSumHandler()
	sumHandler.Getsum(router)
	newHandler := srv.NewHandler(router)
	if err := srv.Handle(newHandler); err != nil {
		logger.Fatal(err)
	}

	// Create service
	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(nacosRegistry),
		//micro.Registry(etcd.NewRegistry()),
	)
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}

}
