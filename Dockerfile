FROM golang:1.15.8-alpine3.13 AS builder
WORKDIR /build
RUN adduser -u 10001 -D app-runner
ENV GOPROXY https://goproxy.cn
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o goapp ./app/main.go

FROM alpine:3.13 AS final
WORKDIR /run
COPY --from=builder /build/goapp /run/
COPY --from=builder /build/config /run/config
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER app-runner

ENV GIN_MODE=release \
    PORT=8099

EXPOSE 8099
ENTRYPOINT ["/run/goapp", "-f" ,"/run/config/example.json"]