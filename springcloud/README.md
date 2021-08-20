# micro-service
微服务技术选型 springcloud、dubbo、go、.net、rust

--server.port=8080
--spring.application.name=provider
--spring.cloud.nacos.config.server-addr=localhost:8848
--spring.cloud.nacos.discovery.server-addr=localhost:8848
--management.endpoints.web.exposure.include=*

--server.port=8081
--spring.application.name=consumer
--spring.cloud.nacos.config.server-addr=localhost:8848
--spring.cloud.nacos.discovery.server-addr=localhost:8848
--management.endpoints.web.exposure.include=*


--server.port=8888
--spring.application.name=gateway
--spring.cloud.nacos.config.server-addr=localhost:8848
--spring.cloud.nacos.discovery.server-addr=localhost:8848
--management.endpoints.web.exposure.include=*
--management.endpoint.gateway.enabled=true
--spring.cloud.gateway.discovery.locator.enabled=true