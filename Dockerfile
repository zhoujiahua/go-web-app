# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.17

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 文件复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将项目源代码复制到工作目录
COPY . .

# 构建项目
RUN go build -o main .

# 暴露端口，确保与 main.go 中设置的端口一致
EXPOSE 8080

# 运行应用程序
CMD ["/app/main"]
