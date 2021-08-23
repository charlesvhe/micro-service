package handler

import (
	"context"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gin-gonic/gin"
	"micro-service/golang/consumer/proto/sum"
	"strconv"
)

type sumHandler struct {
}

func NewSumHandler() *sumHandler {
	return &sumHandler{}
}

func (a *sumHandler) Getsum(router *gin.Engine) {

	router.GET("/getsum/:params", func(c *gin.Context) {
		params := c.Param("params")
		inPut, _ := strconv.ParseInt(params, 10, 64)
		service := micro.NewService()
		service.Init()
		client := sum.NewSumService("provider", service.Client())
		resp, err := client.GetSum(context.Background(), &sum.SumRequest{
			Input: inPut,
		})
		if err != nil {
			logger.Error("无法调用sum-srv服务，请检查sum-srv是否存在")
			c.JSON(500, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": resp.Output})
	})
}
