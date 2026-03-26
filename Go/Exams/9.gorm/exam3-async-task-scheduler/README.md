# 实践题：异步任务调度系统

## 需求描述

你需要实现一个**异步任务调度系统**，该系统接收任务请求，将任务放入队列异步执行，并通过 Channel 实现任务状态的通知机制。

### 核心功能

1. **任务提交接口** (`POST /api/tasks`)
   - 接收任务类型（`email`, `report`, `cleanup`）和任务参数（JSON 格式）
   - 将任务保存到数据库，状态为 `pending`
   - 将任务 ID 推入 Channel，触发异步执行
   - 立即返回任务 ID 和初始状态

2. **任务查询接口** (`GET /api/tasks/:id`)
   - 根据任务 ID 查询任务的当前状态和执行结果

3. **任务列表接口** (`GET /api/tasks`)
   - 支持按状态过滤（`pending`, `running`, `completed`, `failed`）
   - 支持分页（`page`, `page_size`）

4. **任务统计接口** (`GET /api/tasks/stats`)
   - 返回各状态任务的数量统计

5. **异步任务处理器**
   - 使用 Channel 作为任务队列，缓冲待处理的任务 ID
   - 多个 Worker Goroutine 从 Channel 中获取任务 ID 并执行
   - 任务执行时更新数据库状态为 `running`，完成后更新为 `completed` 或 `failed`
   - 记录任务执行结果或错误信息

### 任务执行逻辑

不同类型的任务有不同的模拟执行逻辑：

- **email**: 模拟发送邮件，执行时间 1-3 秒，成功率 90%
- **report**: 模拟生成报表，执行时间 3-5 秒，成功率 80%
- **cleanup**: 模拟清理任务，执行时间 0.5-1 秒，成功率 95%

### 数据库模型

```go
type Task struct {
    ID        uint           `gorm:"primaryKey"`
    Type      string         `gorm:"not null"` // email, report, cleanup
    Params    string         `gorm:"type:json"` // 任务参数 JSON
    Status    string         `gorm:"not null;index"` // pending, running, completed, failed
    Result    string         // 执行结果
    Error     string         // 错误信息
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### 技术要求

1. 使用 **Gin** 框架实现 HTTP API
2. 使用 **GORM** 进行数据库操作（使用 SQLite 内存模式或文件模式）
3. 使用 **Channel** 实现任务队列和 Goroutine 间的通信
4. 实现优雅关闭：接收到退出信号时，停止接收新任务，等待正在执行的任务完成后再退出

### 项目结构

```
exam3-async-task-scheduler/
├── main.go          // 程序入口，初始化数据库、路由、Worker
├── handler.go       // HTTP 处理器
├── model.go         // 数据模型
├── scheduler.go     // 任务调度器（Channel + Worker）
├── go.mod           // 模块定义
└── README.md        // 本文件
```

### 运行要求

```bash
# 初始化模块
go mod init async-task-scheduler

# 添加依赖
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

# 运行
go run .
```

### 测试示例

```bash
# 提交任务
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{"type":"email","params":{"to":"user@example.com","subject":"Hello"}}'

# 查询任务
curl http://localhost:8080/api/tasks/1

# 获取任务列表
curl "http://localhost:8080/api/tasks?status=pending&page=1&page_size=10"

# 获取统计
curl http://localhost:8080/api/tasks/stats
```

### 进阶要求（可选）

1. 使用 `context.Context` 实现任务取消机制
2. 实现任务重试逻辑（失败任务自动重试最多 3 次）
3. 添加中间件记录请求日志
