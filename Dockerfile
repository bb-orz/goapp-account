# 构建阶段的golang镜像
FROM golang:1.14 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN make build-linux


# 运行阶段指定scratch作为基础镜像
FROM scratch as goapp

WORKDIR /app

# 拷贝上一阶段的二进制文件及配置文件
COPY --from=builder /app/goapp .
COPY --from=builder /app/config .

# 如使用https，拷贝证书
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/cert

# 如使用 gin web 框架，可指定运行时环境变量
#ENV GIN_MODE=release \
#    PORT=8090

EXPOSE 8090

ENTRYPOINT ["./goapp -f ./config/example.json"]
