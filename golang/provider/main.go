package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/nacos/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"micro-service/golang/provider/handler"
	"micro-service/golang/provider/proto/sum"
	"os"
)

const (
	defaultNacosAddr = "127.0.0.1:8848"
	serviceName      = "go-provider"
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
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Registry(nacosRegistry),
		micro.Address(":8081"),
	)
	//服务初始化
	service.Init(
		micro.AfterStart(func() error {
			logger.Infof("%v服务启动完成", serviceName)
			return nil
		}),
	)

	sum.RegisterSumHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}

}
