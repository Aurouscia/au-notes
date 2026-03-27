package main

import (
	"time"
)

// TaskStatus 定义任务状态常量
type TaskStatus string

const (
	TaskPending   TaskStatus = "pending"
	TaskRunning   TaskStatus = "running"
	TaskCompleted TaskStatus = "completed"
	TaskFailed    TaskStatus = "failed"
)

// Task 任务模型
// 要求：
// 1. ID 为主键，自增
// 2. Type 字段不为空
// 3. Status 字段不为空，并添加索引
// 4. Params 和 Result 使用合适的数据类型存储 JSON 字符串
type Task struct {
	ID        uint   `gorm:"primary_key;auto_increment"` // ❌ 注意标签的分隔是分号
	Type      string `gorm:"not null"`                   // ❌ not null 中间没有空格
	Params    string
	Status    string `gorm:"not null;index"`
	Result    string
	Error     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TaskRequest 任务请求结构体
type TaskRequest struct {
	Type   string                 `json:"type" binding:"required,oneof=email report cleanup"`
	Params map[string]interface{} `json:"params"`
}

// TaskResponse 任务响应结构体
type TaskResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	Params    string    `json:"params"`
	Status    string    `json:"status"`
	Result    string    `json:"result,omitempty"`
	Error     string    `json:"error,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TaskStats 任务统计
type TaskStats struct {
	Pending   int64 `json:"pending"`
	Running   int64 `json:"running"`
	Completed int64 `json:"completed"`
	Failed    int64 `json:"failed"`
	Total     int64 `json:"total"`
}

// ToResponse 转换 Task 为 TaskResponse
func (t *Task) ToResponse() TaskResponse {
	return TaskResponse{
		ID:        t.ID,
		Type:      t.Type,
		Params:    t.Params,
		Status:    t.Status,
		Result:    t.Result,
		Error:     t.Error,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
