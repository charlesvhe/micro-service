package service

import "github.com/asim/go-micro/v3/logger"

//getSum的具体逻辑,累加

func GetSum(intputs ...int64) int64 {

	var ret int64
	for _, v := range intputs {
		ret += v
	}
	logger.Info("sum-srv被http-client调用了")
	logger.Infof("累加结果为：%v\n", ret)
	return ret

}
