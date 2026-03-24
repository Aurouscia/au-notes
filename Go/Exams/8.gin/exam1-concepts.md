# Gin 基础概念题

## 一、填空题

1. Gin 框架中，创建一个带有默认中间件（Logger 和 Recovery）的引擎使用函数 `________`，创建一个不带任何中间件的纯净引擎使用函数 `________`。
    - gin.Default()
    - gin.New()

2. Gin 中的 `gin.H` 本质上是 `________` 类型的别名。
    - `map[string]any`

3. 在 Handler 函数中，获取 URL 路径参数（如 `/user/:id` 中的 id）使用 `c.________("id")`，获取 URL 查询参数（如 `?name=xxx`）使用 `c.________("name")`。
    - Param
    - Query

4. 在 Gin 中，中间件是一个返回 `________` 类型的函数。
    - void

5. 使用 `c.________(&obj)` 可以将请求体自动绑定到结构体，并进行验证。
    - ShouldBindJson

## 二、判断题（正确填✅，错误填❌）

1. `gin.Default()` 和 `gin.New()` 创建的引擎完全相同，没有任何区别。
    - 错：Default 创建的带有 Logger 和 Recovery 中间件，New 是没有的

2. `c.JSON(200, data)` 会自动设置 `Content-Type: application/json` 响应头。
    - 对

3. 在 Gin 中，路由分组只能用于组织 URL 前缀，不能应用公共中间件。
    - 错

4. `c.Next()` 用于调用后续中间件或 Handler，`c.Abort()` 用于终止后续处理。
    - 对

5. `c.Param("id")` 可以获取 URL 查询参数 `?id=123` 的值。
    - 错，`?id=123`应该通过`c.Query("id")`获取

## 三、简答题

1. 请简述 `gin.Context` 的作用，并列举至少 3 个常用的方法及其用途。
    - gin.Context 用于读取请求参数和写入响应
    - `JSON` 写入 json 格式的响应
    - `Query` 读取查询字符串参数
    - `Param` 读取路径参数

2. 解释 Gin 中全局中间件和局部中间件的区别，并分别给出使用示例。
    - 全局：应用到所有 endpoint 上的中间件
        - `e.Use(middleware)`
    - 局部：应用到特定 endpoint 上的中间件
        - `e.GET("/xxx", middleware, handler)`

3. 以下代码有什么问题？请指出并修正：

    ```go
    func main() {
        r := gin.Default()
        
        r.GET("/user/:id", func(c *gin.Context) {
            id := c.Query("id")
            c.String(200, "User ID: %s", id)
        })
        
        r.Run()
    }
    ```
    - 获取 `/user/:id` 中的 id，应该使用 `c.Param("id")`

4. 请说明 `c.ShouldBindJSON()` 和 `c.BindJSON()` 的区别。
    - ShouldBindJSON() 会尝试绑定，失败了会返回 error 对象
    - BindJSON() 底层是 MustBindJSON() ，失败了会 c.Abort() 终止 handler 并返回 400

5. 在 Gin 中如何实现以下需求：
    - 所有请求都需要记录日志（全局中间件）
    - `/api` 下的路由需要 JWT 认证（分组中间件）
    - `/health` 路由不需要任何认证，直接访问
    ```go
        r := gin.New()
        r.Use(gin.Logger())
        r.Group("/api", Jwt()){
            ...
        }
    ```
