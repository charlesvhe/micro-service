# micro-service
微服务技术选型 springcloud、dubbo、go、.net、rust

# micro-go v4版本有问题，先用v3
https://github.com/asim/go-micro/tree/v3.7.0

# 代码生成 确保 $GOPATH/bin 下有 protoc protoc-gen-go protoc-gen-micro
# provider dir:
protoc --proto_path=../../proto --micro_out=. --go_out=. provider.proto

# consumer dir:
protoc --proto_path=../../proto --micro_out=. --go_out=. provider.proto consumer.proto

# kerya 测试