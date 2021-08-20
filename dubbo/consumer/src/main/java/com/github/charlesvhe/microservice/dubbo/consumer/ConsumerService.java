package com.github.charlesvhe.microservice.dubbo.consumer;

import com.github.charlesvhe.microservice.dubbo.api.ConsumerApi;
import com.github.charlesvhe.microservice.dubbo.api.ProviderApi;
import org.apache.dubbo.config.annotation.DubboReference;
import org.apache.dubbo.config.annotation.DubboService;

@DubboService
public class ConsumerService implements ConsumerApi {
    @DubboReference
    private ProviderApi providerApi;
    @Override
    public String test(String name) {
        return "Consumer " + providerApi.test(name);
    }
}
