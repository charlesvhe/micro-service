# docker build -t charlesvhe/apisix:all-in-one .
# docker push charlesvhe/apisix:all-in-one

docker run -d --name apisix \
-v `pwd`/apisix.yaml:/usr/local/apisix/conf/config.yaml \
-v `pwd`/dashboard.yaml:/usr/local/apisix-dashboard/conf/conf.yaml \
--link nacos \
-p 9080:9080 -p 9443:9443 -p 9000:9000 \
charlesvhe/apisix:all-in-one


# for https://github.com/apache/apisix/issues/5345
docker run -d --name etcd-server \
--env ALLOW_NONE_AUTHENTICATION=yes \
--env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
bitnami/etcd:3.5.1

docker run -it --name apisix \
--link etcd-server \
-v `pwd`/apisix.yaml:/usr/local/apisix/conf/config.yaml \
-p 9080:9080 -p 9443:9443 \
apache/apisix:2.10.0-alpine
