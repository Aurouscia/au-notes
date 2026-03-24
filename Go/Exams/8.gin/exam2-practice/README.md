# 实践题：用户管理系统 API

## 需求

使用 Gin 框架实现一个简单的用户管理系统 REST API。

## 功能要求

### 1. 数据模型

```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Age      int    `json:"age" binding:"gte=0,lte=150"`
}
```

### 2. API 端点

| 方法 | 路径 | 功能 | 说明 |
|------|------|------|------|
| GET | `/users` | 获取所有用户 | 返回用户列表 |
| GET | `/users/:id` | 获取单个用户 | 根据 ID 查找用户 |
| POST | `/users` | 创建用户 | 接收 JSON，自动分配 ID |
| PUT | `/users/:id` | 更新用户 | 根据 ID 更新用户信息 |
| DELETE | `/users/:id` | 删除用户 | 根据 ID 删除用户 |

### 3. 具体要求

1. **数据存储**：使用内存存储（`map[int]User` 或 `[]User`），无需数据库

2. **响应格式**：统一返回 JSON 格式
   - 成功：`{"code": 0, "message": "success", "data": ...}`
   - 失败：`{"code": 1, "message": "错误信息"}`

3. **错误处理**：
   - 用户不存在时返回 404
   - 参数验证失败时返回 400
   - 返回适当的错误信息

4. **中间件**：
   - 添加一个日志中间件，记录每个请求的方法、路径和处理时间

### 4. 测试示例

```bash
# 创建用户
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"username":"张三","email":"zhangsan@example.com","age":25}'

# 获取所有用户
curl http://localhost:8080/users

# 获取单个用户
curl http://localhost:8080/users/1

# 更新用户
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"username":"张三","email":"zhangsan@new.com","age":26}'

# 删除用户
curl -X DELETE http://localhost:8080/users/1
```

## 项目结构

```
exam2-practice/
├── go.mod
├── main.go      # 主程序入口
└── README.md    # 本文件
```

## 提示

1. 使用 `gin.Default()` 创建引擎
2. 使用 `c.ShouldBindJSON()` 绑定请求体
3. 使用 `c.Param("id")` 获取路径参数
4. 使用 `c.JSON()` 返回响应
5. 可以使用全局变量存储用户数据（简化实现）
