# micro-service
微服务技术选型 springcloud、dubbo、go、.net、rust


--spring.application.name=provider
--dubbo.application.name=dubbo-provider
--dubbo.registry.address=nacos://localhost:8848?username=nacos&password=nacos
--dubbo.protocol.name=dubbo
--dubbo.protocol.port=20880

--spring.application.name=consumer
--dubbo.application.name=dubbo-consumer
--dubbo.registry.address=nacos://localhost:8848?username=nacos&password=nacos
--dubbo.protocol.name=dubbo
--dubbo.protocol.port=30880

--server.port=8888
--spring.application.name=bff
--dubbo.application.name=dubbo-bff
--dubbo.registry.address=nacos://localhost:8848?username=nacos&password=nacos
--dubbo.protocol.name=dubbo
--dubbo.protocol.port=40880
--management.endpoints.web.exposure.include=*