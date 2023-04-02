# Go Web App

Go Web App 是一个简单的 Go 语言书籍管理 Web 应用程序。项目使用 [Gin Web Framework](https://github.com/gin-gonic/gin) 进行开发。

## 功能

- 获取所有书籍
- 获取单个书籍
- 添加书籍
- 修改书籍
- 删除书籍

## 安装

确保您的计算机已安装 [Go](https://golang.org/doc/install)。

克隆此存储库：

```bash
git clone https://github.com/zhoujiahua/go-web-app.git
```

进入项目目录：

```
bashCopy code
cd go-web-app
```

获取依赖：

```
bashCopy code
go get
```

## 运行

在项目根目录下运行：

```
bashCopy code
go run main.go
```

应用程序将在 `http://localhost:8080` 上运行。

## API

以下是 API 路由：

### 获取所有书籍

- 路径：`/api/v1/books`
- 方法：`GET`

### 获取单个书籍

- 路径：`/api/v1/books/:id`
- 方法：`GET`

### 添加书籍

- 路径：`/api/v1/books`
- 方法：`POST`
- 请求参数：
  - `title` (string) - 书名
  - `author` (string) - 作者

### 修改书籍

- 路径：`/api/v1/books/:id`
- 方法：`PUT`
- 请求参数：
  - `title` (string) - 书名
  - `author` (string) - 作者

### 删除书籍

- 路径：`/api/v1/books/:id`
- 方法：`DELETE`

## 许可

本项目采用 [Apache License](http://www.apache.org/licenses/) 进行许可。
