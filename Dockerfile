FROM golang:latest as builder
ARG GOPROXY=https://goproxy.cn
COPY . /src
RUN apt-get update && \
    cd /src && \
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/api api.go

FROM ubuntu:18.04

RUN apt-get update && apt-get install gettext-base
RUN mkdir -p /app/etc/
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /src/build/api /app/api
COPY etc/api.template.yaml /app/etc/api.template.yaml

EXPOSE 8888
CMD envsubst < /app/etc/api.template.yaml > /app/etc/api.yaml && ./api