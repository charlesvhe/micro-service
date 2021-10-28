ARG ETCD_VERSION=v3.4.14
ARG APISIX_VERSION=2.10.0-alpine
ARG APISIX_DASHBOARD_VERSION=2.9.0

# build etcd
FROM alpine:3.13 AS etcd-stage
ARG ETCD_VERSION

WORKDIR /tmp
RUN wget https://github.com/etcd-io/etcd/releases/download/${ETCD_VERSION}/etcd-${ETCD_VERSION}-linux-amd64.tar.gz \
    && tar -zxvf etcd-${ETCD_VERSION}-linux-amd64.tar.gz \
    && ln -s etcd-${ETCD_VERSION}-linux-amd64 etcd

# use apisix-dashboard
FROM apache/apisix-dashboard:${APISIX_DASHBOARD_VERSION} AS apisix-dashboard-stage

# combine all the resources into apisix-dashboard
FROM apache/apisix:${APISIX_VERSION} AS apisix-stage

ARG ETCD_VERSION
ARG APISIX_VERSION
ARG APISIX_DASHBOARD_VERSION

LABEL etcd_version="${ETCD_VERSION}"
LABEL apisix_version="${APISIX_VERSION}"
LABEL apisix_dashboard_version="${APISIX_DASHBOARD_VERSION}"

COPY --from=etcd-stage /tmp/etcd/etcd /usr/bin/etcd
COPY --from=etcd-stage /tmp/etcd/etcdctl /usr/bin/etcdctl

COPY --from=apisix-dashboard-stage /usr/local/apisix-dashboard /usr/local/apisix-dashboard

ENV PATH=$PATH:/usr/local/openresty/luajit/bin:/usr/local/openresty/nginx/sbin:/usr/local/openresty/bin

EXPOSE 9080 9443 2379 2380 9000

CMD ["sh", "-c", "(nohup etcd >/tmp/etcd.log 2>&1 &) && sleep 10 && (cd /usr/local/apisix-dashboard && /usr/local/apisix-dashboard/manager-api &) && cd /usr/local/apisix && /usr/bin/apisix init && /usr/bin/apisix init_etcd && /usr/local/openresty/bin/openresty -p /usr/local/apisix -g 'daemon off;'"]

STOPSIGNAL SIGQUIT