# Gin Web 框架基础

## 什么是 Gin？

Gin 是一个用 Go 语言编写的高性能 HTTP Web 框架，它具有以下特点：

- **高性能**：基于 httprouter，路由匹配速度比传统框架快近 40 倍
- **轻量级**：设计简洁，内存占用小，适合高并发场景
- **易用性**：API 设计友好，类似 Express.js 的风格
- **功能丰富**：内置中间件、JSON 验证、路由分组等功能

## 安装

```bash
go get -u github.com/gin-gonic/gin
```

## 核心概念

### 1. Engine（引擎）

`gin.Engine` 是整个框架的核心，负责路由注册和中间件管理：

```go
// 创建一个默认的 Engine（自带 Logger 和 Recovery 中间件）
r := gin.Default()

// 或者创建一个纯净的 Engine（不带任何中间件）
r := gin.New()
```

### 2. 路由注册

Gin 支持所有 HTTP 方法：

```go
r.GET("/path", handler)
r.POST("/path", handler)
r.PUT("/path", handler)
r.DELETE("/path", handler)
r.PATCH("/path", handler)
r.HEAD("/path", handler)
r.OPTIONS("/path", handler)

// 支持任意方法
r.Any("/path", handler)
```

### 3. Handler 函数

Handler 是处理请求的核心函数，签名如下：

```go
func handler(c *gin.Context) {
    // c 是 *gin.Context，封装了请求和响应
}
```

### 4. gin.Context

`gin.Context` 是处理 HTTP 请求的核心类型，提供了丰富的 API：

**获取请求信息：**
- `c.Param("name")` - 获取 URL 路径参数
- `c.Query("name")` - 获取 URL 查询参数
- `c.PostForm("name")` - 获取表单数据
- `c.Bind(&obj)` - 绑定并验证请求体到结构体

**返回响应：**
- `c.String(code, "text")` - 返回纯文本
- `c.JSON(code, obj)` - 返回 JSON
- `c.XML(code, obj)` - 返回 XML
- `c.HTML(code, "template", data)` - 渲染 HTML 模板

## 快速开始

### 第一个 Gin 应用

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // 创建默认路由引擎
    r := gin.Default()
    
    // 定义路由
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
    
    // 启动服务，默认监听 :8080
    r.Run()
}
```

运行后访问 `http://localhost:8080/ping`，会返回 `{"message":"pong"}`。

### gin.H

`gin.H` 就是 `map[string]any` 的别名，可以在不定义 struct 的情况下返回 json 数据（仅 prototype 或简单情况用）

### 路由参数

```go
// 路径参数
r.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
})

// 查询参数
r.GET("/welcome", func(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname") // 等同于 c.Request.URL.Query().Get("lastname")
    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
})
```

### 表单数据

```go
r.POST("/form", func(c *gin.Context) {
    message := c.PostForm("message")
    nick := c.DefaultPostForm("nick", "anonymous")
    c.JSON(http.StatusOK, gin.H{
        "message": message,
        "nick":    nick,
    })
})
```

### JSON 绑定

```go
type Login struct {
    User     string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required"`
}

r.POST("/login", func(c *gin.Context) {
    var json Login
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if json.User != "admin" || json.Password != "123456" {
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
})
```

## 路由分组

路由分组可以将相关路由组织在一起，并应用公共中间件：

```go
func main() {
    r := gin.Default()
    
    // 简单分组（打括号是在构造 group 结构体）
    v1 := r.Group("/v1")
    {
        v1.GET("/login", loginEndpoint)
        v1.GET("/submit", submitEndpoint)
        v1.GET("/read", readEndpoint)
    }
    
    // 带中间件的分组
    authorized := r.Group("/admin", AuthRequired())
    {
        authorized.POST("/users", createUser)
        authorized.GET("/users", getUsers)
    }
    
    r.Run()
}
```

## 中间件

### 什么是中间件？

中间件是位于请求和 Handler 之间的函数，可以在请求处理前后执行逻辑。

### 全局中间件

```go
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 请求前
        t := time.Now()
        
        // 继续执行后续 Handler
        c.Next()
        
        // 请求后
        latency := time.Since(t)
        log.Printf("%s %s %s", c.Request.Method, c.Request.URL, latency)
    }
}

func main() {
    r := gin.New()
    r.Use(Logger()) // 注册全局中间件
    // ...
}
```

### 局部中间件

```go
// 单个路由
r.GET("/path", AuthRequired(), handler)

// 路由组
authorized := r.Group("/admin")
authorized.Use(AuthRequired())
{
    authorized.GET("/secrets", getSecrets)
}
```

### 常用内置中间件

```go
// Logger - 日志记录
r.Use(gin.Logger())

// Recovery - 从 panic 恢复，防止服务器崩溃
r.Use(gin.Recovery())

// BasicAuth - HTTP 基本认证
authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
    "admin": "password123",
}))
```

## 静态文件服务

```go
// 提供静态文件
r.Static("/assets", "./assets")

// 提供单个静态文件
r.StaticFile("/favicon.ico", "./resources/favicon.ico")

// 加载 HTML 模板
r.LoadHTMLGlob("templates/*")
r.LoadHTMLFiles("templates/index.html", "templates/login.html")
```

## 总结

| 概念 | 说明 |
|------|------|
| Engine | Web 服务器核心，负责路由和中间件 |
| gin.Context | 请求上下文，封装 Request 和 Response |
| HandlerFunc | 处理请求的函数签名 `func(*gin.Context)` |
| 路由分组 | 使用 `Group()` 组织相关路由 |
| 中间件 | 在请求处理前后执行逻辑的函数 |

掌握以上内容后，你就可以使用 Gin 开发基本的 Web 服务和 REST API 了。
