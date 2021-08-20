package com.github.charlesvhe.microservice.dubbo.bff;

import com.github.charlesvhe.microservice.dubbo.api.ConsumerApi;
import com.github.charlesvhe.microservice.dubbo.api.ProviderApi;
import org.apache.dubbo.config.annotation.DubboReference;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class BffController {
    @DubboReference
    private ProviderApi providerApi;
    @DubboReference
    private ConsumerApi consumerApi;

    @GetMapping("/testProvider")
    public String testProvider(String name){
        return providerApi.test(name);
    }
    @GetMapping("/testConsumer")
    public String testConsumer(String name){
        return consumerApi.test(name);
    }
}
