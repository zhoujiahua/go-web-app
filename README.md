# Go Web App

Go Web App 是一个简单的 Go 语言书籍管理 Web 应用程序。项目使用 [Gin Web Framework](https://github.com/gin-gonic/gin) 进行开发，并遵循 Apache License 2.0 许可。

### 目录结构

> 项目的目录结构说明：

```html
go-web-app 
│
├── controllers     # 存放所有的控制器，用于处理 HTTP 请求
│   ├── book_controller.go
│
├── models          # 存放所有的数据模型，用于处理业务逻辑
│   ├── book.go
│
├── routes          # 存放路由配置，用于定义 API 路径和对应的处理函数
│   ├── book_router.go
│
├── utils           # 存放工具类，例如：自定义的 JSON 响应函数
│   ├── response.go
│
├── main.go         # 应用程序的入口
├── README.md       # 项目说明文档
└── go.mod          # Go 项目的依赖管理文件
```

### 功能API

| API名称      | 路由                | 方法   | 请求参数                                            |
| ------------ | ------------------- | ------ | --------------------------------------------------- |
| 获取所有书籍 | `/api/v1/books`     | GET    | 无                                                  |
| 获取单个书籍 | `/api/v1/books/:id` | GET    | 无                                                  |
| 添加书籍     | `/api/v1/books`     | POST   | `title` (string) - 书名<br>`author` (string) - 作者 |
| 修改书籍     | `/api/v1/books/:id` | PUT    | `title` (string) - 书名<br>`author` (string) - 作者 |
| 删除书籍     | `/api/v1/books/:id` | DELETE | 无                                                  |

### 安装

确保您的计算机已安装 [Go](https://golang.org/doc/install)。

克隆此存储库：

```bash
git clone https://github.com/zhoujiahua/go-web-app.git
```

进入项目目录：

```
cd go-web-app
```

获取依赖：

```bash
go get
```

### 运行

在项目根目录下运行：

```bash
go run main.go
```

应用程序将在 `http://localhost:9527` 上运行。

### Docker

> 打开终端或命令提示符，运行以下命令构建 Docker 镜像：

```bash
# 请注意，这里的 `go-web-app` 是镜像的名称，你可以根据自己的需要更改它。
docker build -t go-web-app .
```

> 镜像构建完成后，运行以下命令启动容器：

```bash
# 同样地，`go-web-app-container` 是容器的名称，你可以根据自己的需要更改它。
docker run -d -p 9527:9527 --name go-web-app-container go-web-app
```

现在，你已经成功地将项目构建成 Docker 镜像并运行了容器。在浏览器中访问 `http://localhost:9527/api/v1/books`，你应该能看到项目正常运行并返回 JSON 数据。

### Docker-Compose

> 打开终端或命令提示符，运行以下命令启动 Docker Compose：

```bash
# 启动
docker-compose up --build

# 停止
docker-compose down
```

现在，你已经成功地使用 Docker Compose 构建并运行了该项目。在浏览器中访问 `http://localhost:9527/api/v1/books`，你应该能看到项目正常运行并返回 JSON 数据。

### 许可

本项目采用 [Apache License](http://www.apache.org/licenses/) 进行许可。
