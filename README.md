# micro-service
微服务技术选型 springcloud、dubbo、go、.net、rust

映射http和grpc端口
docker run -d --name nacos -e MODE=standalone -p 8848:8848 -p 9848:9848 nacos/nacos-server

http://localhost:8848/nacos/index.html

docker run -d --name apisix -v `pwd`/apisix/apisix.yaml:/usr/local/apisix/conf/config.yaml \
    -v `pwd`/apisix/dashboard.yaml:/usr/local/apisix-dashboard/conf/conf.yaml \
    --link nacos:nacos -p 9080:9080 -p 9091:9091 -p 2379:2379 -p 9000:9000 charlesvhe/apisix:all-in-one
    
http://localhost:9000




