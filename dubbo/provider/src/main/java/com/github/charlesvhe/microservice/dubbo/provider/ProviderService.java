package com.github.charlesvhe.microservice.dubbo.provider;

import com.github.charlesvhe.microservice.dubbo.api.ProviderApi;
import org.apache.dubbo.config.annotation.DubboService;

@DubboService
public class ProviderService implements ProviderApi {
    @Override
    public String test(String name) {
        return "Hello " + name + "! " + System.currentTimeMillis();
    }
}
