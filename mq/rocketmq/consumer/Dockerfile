FROM golang:latest as builder

WORKDIR ${GOPATH}/src/learngo/rocketmq/consumer

COPY / .

# 每次构建程序时都会创建一层新镜像，名称是 none
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM centos

WORKDIR /data

COPY --from=builder /go/src/learngo/rocketmq/consumer .