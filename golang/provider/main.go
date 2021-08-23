package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/isfk/go-micro-plugins/registry/nacos/v3"
	"micro-service/golang/provider/handler"
	"micro-service/golang/provider/proto/sum"
	"os"
)

var etcdReg registry.Registry

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
		options.Context = context.WithValue(context.Background(), &nacos.NacosNamespaceContextKey{}, nacosNamespace)

	})
	service := micro.NewService(
		micro.Name("provider"),
		micro.Registry(nacosRegistry),
		micro.Address(":8081"),
	)
	//服务初始化
	service.Init(
		micro.BeforeStart(func() error {
			logger.Info("provider服务启动前日志")
			return nil
		}),
		micro.AfterStart(func() error {
			logger.Infof("provider服务启动,注册地址为 %v\n", nacosAddr)
			return nil
		}),
	)

	sum.RegisterSumHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}

}
