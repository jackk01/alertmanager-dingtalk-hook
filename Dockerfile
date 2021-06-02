FROM golang:alpine AS build

ENV GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o dingtalk

FROM alpine:latest AS prod

RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories && \
    apk update && \
    apk --no-cache add ca-certificates tzdata && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

COPY --from=build /app/dingtalk /usr/bin/dingtalk

EXPOSE 5000

ENTRYPOINT ["dingtalk"]