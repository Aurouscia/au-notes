package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TaskHandler HTTP 处理器
type TaskHandler struct {
	db        *gorm.DB
	scheduler *TaskScheduler
}

// NewTaskHandler 创建处理器
func NewTaskHandler(db *gorm.DB, scheduler *TaskScheduler) *TaskHandler {
	return &TaskHandler{
		db:        db,
		scheduler: scheduler,
	}
}

// CreateTask 创建任务
// TODO: 实现任务提交接口
// POST /api/tasks
// 要求：
// 1. 绑定请求参数（type, params）
// 2. 将任务保存到数据库，初始状态为 pending
// 3. 将任务 ID 提交到调度器队列
// 4. 返回任务 ID 和初始状态
func (h *TaskHandler) CreateTask(c *gin.Context) {
	// 请在此处实现
}

// GetTask 获取任务详情
// TODO: 实现任务查询接口
// GET /api/tasks/:id
// 要求：
// 1. 从 URL 参数获取任务 ID
// 2. 从数据库查询任务
// 3. 如果任务不存在返回 404
// 4. 返回任务详情
func (h *TaskHandler) GetTask(c *gin.Context) {
	// 请在此处实现
}

// ListTasks 获取任务列表
// TODO: 实现任务列表接口
// GET /api/tasks?status=&page=&page_size=
// 要求：
// 1. 获取查询参数：status（可选）、page（默认1）、page_size（默认10，最大100）
// 2. 根据 status 过滤（如果提供）
// 3. 实现分页查询
// 4. 返回任务列表和总数
func (h *TaskHandler) ListTasks(c *gin.Context) {
	// 请在此处实现
}

// GetTaskStats 获取任务统计
// TODO: 实现任务统计接口
// GET /api/tasks/stats
// 要求：
// 1. 使用 GORM 的 Count 方法统计各状态任务数量
// 2. 返回统计结果
func (h *TaskHandler) GetTaskStats(c *gin.Context) {
	// 请在此处实现
}

// paramsToString 将 map 转换为 JSON 字符串（辅助函数）
func paramsToString(params map[string]interface{}) string {
	if params == nil {
		return "{}"
	}
	data, _ := json.Marshal(params)
	return string(data)
}
