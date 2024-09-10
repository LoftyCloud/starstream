# 构建阶段
FROM golang:alpine AS builder

# 环境变量
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct

# 设置阿里云镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && apk add --no-cache git

WORKDIR /build

# 复制 go.mod 和 go.sum 并下载依赖
COPY go.mod .
COPY go.sum .
RUN go mod download

# 复制源码
COPY . .

# 列出文件，调试信息
RUN ls -la

# 进行构建
RUN go build -o main

# 运行阶段
FROM scratch
WORKDIR /app
COPY --from=builder /build/main /app

CMD ["./main"]
