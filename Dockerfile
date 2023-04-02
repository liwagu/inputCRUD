# 基于 golang 镜像构建
FROM golang:alpine as builder

# 设置工作目录
WORKDIR /app

# 复制源代码到工作目录
COPY . .

# 编译项目
RUN go build -o inputcrud .

# 基于 alpine 镜像构建
FROM alpine

# 安装数据库客户端
RUN apk add --no-cache mysql-client

# 设置工作目录
WORKDIR /app

# 从 builder 镜像中复制二进制文件
COPY --from=builder /app/inputcrud .

# 设置环境变量
ENV DB_USER=username
ENV DB_PASS=password
ENV DB_HOST=database
ENV DB_NAME=inputcrud

# 暴露端口
EXPOSE 8080

# 运行二进制文件
CMD ["./inputcrud"]

ENV GOPROXY=https://goproxy.cn


