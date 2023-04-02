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

### 功能

- 获取所有书籍
- 获取单个书籍
- 添加书籍
- 修改书籍
- 删除书籍

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

应用程序将在 `http://localhost:8080` 上运行。

### API

以下是 API 路由：

#### 获取所有书籍

- 路径：`/api/v1/books`
- 方法：`GET`

#### 获取单个书籍

- 路径：`/api/v1/books/:id`
- 方法：`GET`

#### 添加书籍

- 路径：`/api/v1/books`
- 方法：`POST`
- 请求参数：
  - `title` (string) - 书名
  - `author` (string) - 作者

#### 修改书籍

- 路径：`/api/v1/books/:id`
- 方法：`PUT`
- 请求参数：
  - `title` (string) - 书名
  - `author` (string) - 作者

#### 删除书籍

- 路径：`/api/v1/books/:id`
- 方法：`DELETE`

### 许可

本项目采用 [Apache License](http://www.apache.org/licenses/) 进行许可。
