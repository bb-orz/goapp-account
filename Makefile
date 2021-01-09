PROJECT="goapp"
MAIN_PATH="./app/main.go"
VERSION="v1.0.0"
DATE= `date +%FT%T%z`

version:
	@echo ${VERSION}

.PHONY: build
build:
	@echo version: ${VERSION} date: ${DATE} os: Mac OS
	@go  build -o ${PROJECT} ${MAIN_PATH}

install:
	@echo download package
	@go mod download

# 交叉编译运行在linux系统环境
build-linux:
	@echo version: ${VERSION} date: ${DATE} os: linux-centOS
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${PROJECT} ${MAIN_PATH}

run:   build
	@./${PROJECT} -f "./config/example.json"


