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

const (
	defaultNacosAddr = "127.0.0.1:8848"
	serviceName      = "go-consumer"
)

func main() {
	// 从环境变量中获取nacos的ip和port
	var nacosAddr string
	nacosAddr = os.Getenv("NacosAddr")
	if nacosAddr == "" {
		nacosAddr = defaultNacosAddr
	}

	nacosRegistry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{nacosAddr}
		options.Context = context.Background()
	})
	srv := httpServer.NewServer(
		server.Name(serviceName),
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
	)
	service.Init(
		micro.AfterStart(func() error {
			logger.Infof("%v 服务启动后日志", serviceName)
			return nil
		}))

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}

}
