# 使用 golang 官方镜像作为基础镜像
FROM golang:1.16-alpine

# 设置工作目录
WORKDIR /app

# 复制 Go 代码到容器中
COPY . .

# 构建应用程序
RUN go build -o go-web-app

# 设置容器入口点
ENTRYPOINT ["./go-web-app"]
