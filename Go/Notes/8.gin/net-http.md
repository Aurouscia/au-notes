# net/http 包详解

Go 标准库中的 `net/http` 包提供了 HTTP 客户端和服务端的基础实现。Gin 框架底层就是基于 `net/http` 构建的。

## 1. 创建 HTTP 服务端

### 最基础的服务器

```go
package main

import (
    "fmt"
    "net/http"
)

// Handler 函数签名：func(w http.ResponseWriter, r *http.Request)
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // 注册路由
    http.HandleFunc("/hello", helloHandler)
    
    // 启动服务器，监听 :8080
    fmt.Println("Server starting on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Server error:", err)
    }
}
```

> 注意：`http.ResponseWriter`不需要星号，这是由于它是接口，不是结构体（和 `context.Context` 不需要星号一个道理  
> 接口变量本身就包含一个指针，指向实际的数据，所以传递接口时，已经在传递引用了

### http.HandleFunc 详解

```go
// 函数签名
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// pattern: URL 路径模式
// handler: 处理函数，接收两个参数：
//   - ResponseWriter: 用于写入响应
//   - *Request: 包含请求的所有信息
```

## 2. ResponseWriter 和 Request

### http.ResponseWriter

用于构造 HTTP 响应：

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 设置状态码（必须在 WriteHeader 之前设置）
    w.WriteHeader(http.StatusOK)  // 200
    
    // 设置响应头
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Custom-Header", "value")
    
    // 写入响应体
    w.Write([]byte("Hello"))
    
    // 或者使用 fmt.Fprintf
    fmt.Fprintf(w, "Hello %s", "World")
}
```

**⚠️ 注意**：`WriteHeader` 只能调用一次，且必须在 `Write` 之前。如果先调用 `Write`，会自动发送 200 状态码。

### http.Request

包含 HTTP 请求的所有信息：

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 请求方法
    method := r.Method  // GET, POST, PUT, DELETE 等
    
    // URL 路径
    path := r.URL.Path  // /hello
    
    // URL 查询参数
    query := r.URL.Query()
    name := query.Get("name")  // ?name=xxx
    
    // 请求头
    userAgent := r.Header.Get("User-Agent")
    
    // 读取请求体
    body, _ := io.ReadAll(r.Body)
    defer r.Body.Close()
}
```

## 3. 获取请求参数

### URL 查询参数

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // URL: /search?q=golang&page=1
    
    query := r.URL.Query()
    
    q := query.Get("q")        // "golang"
    page := query.Get("page")  // "1"
    
    // 获取多个值
    tags := query["tag"]  // ?tag=go&tag=web → ["go", "web"]
}
```

### 表单数据（POST）

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 解析表单数据（必须调用）
    r.ParseForm()
    
    // 获取表单值
    username := r.FormValue("username")
    password := r.Form.Get("password")  // r.Form 是 map[string][]string
    
    // 或者直接用 PostFormValue（只从 POST body 获取）
    email := r.PostFormValue("email")
}
```

### 路径参数（需要额外处理）

标准库不直接支持 `/user/:id` 这种路径参数，需要手动解析：

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // URL: /user/123
    path := r.URL.Path  // "/user/123"
    
    // 手动分割
    parts := strings.Split(path, "/")
    // parts = ["", "user", "123"]
    if len(parts) >= 3 {
        id := parts[2]
        fmt.Fprintf(w, "User ID: %s", id)
    }
}
```

## 4. 返回 JSON 响应

```go
func jsonHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
        "name": "张三",
        "age":  25,
    }
    
    // 设置 Content-Type
    w.Header().Set("Content-Type", "application/json")
    
    // 编码为 JSON 并写入
    json.NewEncoder(w).Encode(data)
}
```

## 5. 使用 http.ServeMux（路由多路复用器）

```go
func main() {
    // 创建自定义的 ServeMux
    mux := http.NewServeMux()
    
    // 注册路由
    mux.HandleFunc("/", homeHandler)
    mux.HandleFunc("/api/users", usersHandler)
    mux.HandleFunc("/api/posts", postsHandler)
    
    // 使用自定义 mux 启动服务器
    http.ListenAndServe(":8080", mux)
}
```

## 6. HTTP 客户端

### 发送 GET 请求

```go
func main() {
    resp, err := http.Get("https://api.github.com/users/github")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    // 读取响应
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```

### 发送 POST 请求

```go
func main() {
    // POST 表单数据
    data := url.Values{}
    data.Set("name", "张三")
    data.Set("age", "25")
    
    resp, err := http.PostForm("https://httpbin.org/post", data)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```

### 发送 JSON POST 请求

```go
func main() {
    jsonData := map[string]string{"name": "张三"}
    jsonBytes, _ := json.Marshal(jsonData)
    
    resp, err := http.Post(
        "https://httpbin.org/post",
        "application/json",
        bytes.NewBuffer(jsonBytes),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```

### 使用 http.Client（更灵活）

```go
func main() {
    // 创建自定义客户端
    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    
    // 创建请求
    req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // 设置请求头
    req.Header.Set("Authorization", "Bearer token123")
    req.Header.Set("Accept", "application/json")
    
    // 发送请求
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    // 处理响应...
}
```

## 7. 处理不同 HTTP 方法

```go
func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handleGet(w, r)
    case http.MethodPost:
        handlePost(w, r)
    case http.MethodPut:
        handlePut(w, r)
    case http.MethodDelete:
        handleDelete(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "Method not allowed: %s", r.Method)
    }
}
```

## 8. 常用状态码常量

```go
http.StatusOK                  // 200
http.StatusCreated             // 201
http.StatusBadRequest          // 400
http.StatusUnauthorized        // 401
http.StatusForbidden           // 403
http.StatusNotFound            // 404
http.StatusMethodNotAllowed    // 405
http.StatusInternalServerError // 500
```

## 9. 与 Gin 的关系

| 特性 | net/http | Gin |
|------|----------|-----|
| 路由参数 | ❌ 不支持 `:id` | ✅ 支持 `/user/:id` |
| 路由分组 | ❌ 不支持 | ✅ 支持 `Group()` |
| 中间件 | ⚠️ 需手动实现 | ✅ 内置支持 |
| JSON 绑定 | ❌ 需手动解析 | ✅ 自动绑定验证 |
| 性能 | 基础 | 高性能（httprouter）|
| 依赖 | 标准库 | 第三方包 |

**总结**：`net/http` 是 Go 标准库的基础实现，适合简单的 HTTP 服务；Gin 在其之上提供了更丰富的功能和更好的性能，适合生产环境的 Web 开发。

## 10. 完整示例：REST API

```go
package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users = []User{
    {ID: 1, Name: "张三"},
    {ID: 2, Name: "李四"},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(users)
    case http.MethodPost:
        var user User
        json.NewDecoder(r.Body).Decode(&user)
        user.ID = len(users) + 1
        users = append(users, user)
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // 从路径提取 ID: /users/1
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    id, _ := strconv.Atoi(parts[2])
    
    switch r.Method {
    case http.MethodGet:
        for _, u := range users {
            if u.ID == id {
                json.NewEncoder(w).Encode(u)
                return
            }
        }
        w.WriteHeader(http.StatusNotFound)
    case http.MethodDelete:
        for i, u := range users {
            if u.ID == id {
                users = append(users[:i], users[i+1:]...)
                w.WriteHeader(http.StatusNoContent)
                return
            }
        }
        w.WriteHeader(http.StatusNotFound)
    }
}

func main() {
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/users/", userHandler)
    
    http.ListenAndServe(":8080", nil)
}
```
