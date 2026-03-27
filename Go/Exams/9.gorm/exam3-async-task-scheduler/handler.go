package main

import (
	"encoding/json"
	"errors"
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
// POST /api/tasks
// 要求：
// 1. 绑定请求参数（type, params）
// 2. 将任务保存到数据库，初始状态为 pending
// 3. 将任务 ID 提交到调度器队列
// 4. 返回任务 ID 和初始状态
func (h *TaskHandler) CreateTask(c *gin.Context) {
	taskReq := TaskRequest{}
	err := c.BindJSON(&taskReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := Task{
		ID:     0,
		Params: paramsToString(taskReq.Params),
		Type:   taskReq.Type,
		Status: "pending",
	}
	err = h.db.Create(&task).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.scheduler.SubmitTask(task.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// GetTask 获取任务详情
// GET /api/tasks/:id
// 要求：
// 1. 从 URL 参数获取任务 ID
// 2. 从数据库查询任务
// 3. 如果任务不存在返回 404
// 4. 返回任务详情
func (h *TaskHandler) GetTask(c *gin.Context) {
	task := Task{}
	err := h.db.First(&task, c.Param("id")).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// ListTasks 获取任务列表
// GET /api/tasks?status=&page=&page_size=
// 要求：
// 1. 获取查询参数：status（可选）、page（默认1）、page_size（默认10，最大100）
// 2. 根据 status 过滤（如果提供）
// 3. 实现分页查询
// 4. 返回任务列表和总数
func (h *TaskHandler) ListTasks(c *gin.Context) {
	tasks := []Task{}
	status := c.Query("status")
	page := c.Query("page")
	pageSize := c.Query("page_size")

	pageNum, convErr := strconv.Atoi(page)
	if convErr != nil {
		pageNum = 1
	}
	pageIndex := pageNum - 1
	pageSizeNum, convErr := strconv.Atoi(pageSize)
	if convErr != nil {
		pageSizeNum = 10
	}

	total :=
		h.db.Model(&Task{}).Count()

	q := h.db.Model(&Task{})
	if status != "" {
		q = q.Where("status = ?", status)
	}
	q = q.Offset(pageIndex * pageSizeNum).Limit(pageSizeNum)
	err := q.Find(&tasks).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTaskStats 获取任务统计
// GET /api/tasks/stats
// 要求：
// 1. 使用 GORM 的 Count 方法统计各状态任务数量
// 2. 返回统计结果
func (h *TaskHandler) GetTaskStats(c *gin.Context) {
	var results any
	err := h.db.Model(&Task{}).Select("Status, count(*) as Total").Group("Status").Find(&results).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, results)
}

// paramsToString 将 map 转换为 JSON 字符串（辅助函数）
func paramsToString(params map[string]interface{}) string {
	if params == nil {
		return "{}"
	}
	data, _ := json.Marshal(params)
	return string(data)
}
