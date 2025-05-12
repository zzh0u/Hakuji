# 使用官方 Go 作为基础镜像
FROM golang:1.23.2

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将应用源代码复制到工作目录
COPY . .

# 编译应用
RUN go build -o main .

# 暴露应用运行的端口
EXPOSE 8000

# 运行编译好的可执行文件
CMD ["./main"]