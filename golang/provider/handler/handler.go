package handler

import (
	"context"
	"micro-service/golang/provider/proto/sum"
	"micro-service/golang/provider/service"
)

// 私有
type handler struct {
}

func (h handler) GetSum(ctx context.Context, request *sum.SumRequest, response *sum.SumResponse) error {

	inputs := make([]int64, 0)
	var i int64
	for i = 0; i <= request.Input; i++ {
		inputs = append(inputs, i)
	}
	response.Output = service.GetSum(inputs...)

	return nil
}

//外部暴露handler
func Handler() sum.SumHandler {
	return handler{}

}
