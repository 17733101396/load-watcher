FROM golang:1.15.5

WORKDIR /go/src/github.com/paypal/load-watcher
COPY . .
RUN make build

FROM alpine:3.12

COPY --from=0 /go/src/github.com/paypal/load-watcher/bin/load-watcher /bin/load-watcher
RUN echo "http://mirrors.aliyun.com/alpine/v3.11/main" > /etc/apk/repositories \
&& apk add --no-cache curl tzdata bash\
&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone \
&& apk del tzdata \
&& rm -rf /var/cache/apk/* \
&& rm -rf /root/.cache \
&& rm -rf /tmp/*
CMD ["/bin/load-watcher"]
