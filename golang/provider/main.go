package main

import (
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
	defaultNacosAddr = "172.30.8.225:8848"
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
			logger.Info("provider服务启动后日志")
			return nil
		}),
	)

	sum.RegisterSumHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}

}
