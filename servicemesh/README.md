# micro-service
微服务技术选型 springcloud、dubbo、go、.net、rust

# 安装apisix网关
kubectl -n default create configmap apisix-config --from-file=../apisix/apisix.yaml
kubectl -n default create configmap apisix-dashboard-config --from-file=../apisix/dashboard.yaml

kubectl -n default apply -f apisix.yaml
kubectl -n default delete -f apisix.yaml

# 安装http-bin
# docker run -d --name httpbin -v `pwd`/jackal.jpg:/usr/local/lib/python3.6/dist-packages/httpbin/templates/images/jackal.jpg -p 8080:80 kennethreitz/httpbin

kubectl -n default apply -f httpbin.yaml
kubectl -n default delete -f httpbin.yaml

# 安装http-bin v2
kubectl -n default create configmap jackal-jpg --from-file=jackal.jpg
kubectl -n default apply -f httpbin2.yaml
kubectl -n default delete -f httpbin2.yaml

# 配置路由 路径改写 正则改写 ^/httpbin2/(.*)     /$1    上游主机 httpbin.default

# 安装istio
https://istio.io/latest/docs/setup/getting-started/
# 配置翻墙
export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890

istioctl install --set profile=minimal -y
kubectl create namespace istio-test
kubectl label namespace istio-test istio-injection=enabled

kubectl get ns --show-labels=true
