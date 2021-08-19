package com.github.charlesvhe.microservice.springcloud.provider;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.context.annotation.Bean;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@EnableDiscoveryClient
@SpringBootApplication
public class ProviderApplication {

    @LoadBalanced
    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }

    @RestController
    public class TestController {
        @GetMapping("/test")
        public String test(String name) {
            return "Hello " + name + "! " + System.currentTimeMillis();
        }
    }

    public static void main(String[] args) {
        SpringApplication.run(ProviderApplication.class, args);
    }
}
