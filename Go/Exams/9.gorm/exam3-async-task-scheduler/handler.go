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
		// ⚠️ 最佳实践：ID 为 0 时 GORM 会自动处理为自增，显式设置 ID 为 0 是多余的（默认就是 0 ）
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
		// ⚠️ 最佳实践：任务提交到队列失败时，应该考虑更新数据库中的任务状态为 failed
		// 或者提供重试机制，避免任务处于 pending 状态但永远不会被执行
		h.db.Model(&task).Update("Status", "failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ❌ 返回响应时建议使用专门的 Response 结构体，而不是直接返回数据库模型
	// 这样可以集中控制返回的字段，避免暴露内部字段
	c.JSON(http.StatusOK, task.ToResponse())
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

	// ⚠️ 最佳实践：page_size 应该限制最大值，防止请求过大导致内存问题
	const maxPageSize = 100
	if pageSizeNum > maxPageSize {
		pageSizeNum = maxPageSize
	}
	if pageSizeNum <= 0 {
		pageSizeNum = 10
	}
	if pageNum <= 0 {
		pageNum = 1
	}

	// ❌ 这里的 total 统计应该在过滤条件之后，否则统计的是全部任务数
	// 如果需要返回符合过滤条件的总数，应该将过滤逻辑应用到 Count 查询
	var total int64 = 0
	countQuery := h.db.Model(&Task{})
	if status != "" {
		countQuery = countQuery.Where("status = ?", status)
	}
	countQuery.Count(&total)

	q := h.db.Model(&Task{})
	if status != "" {
		q = q.Where("status = ?", status)
	}
	q = q.Offset(pageIndex * pageSizeNum).Limit(pageSizeNum)
	err := q.Find(&tasks).Error
	if err != nil {
		// ❌ 错误响应应该使用 gin.H 包装，保持响应格式一致
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ⚠️ 最佳实践：返回响应时转换为 Response 结构体列表
	var responses []TaskResponse
	for _, t := range tasks {
		responses = append(responses, t.ToResponse())
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"list":  responses,
	})
}

// GetTaskStats 获取任务统计
// GET /api/tasks/stats
// 要求：
// 1. 使用 GORM 的 Count 方法统计各状态任务数量
// 2. 返回统计结果
func (h *TaskHandler) GetTaskStats(c *gin.Context) {
	// ❌ 使用 any 类型无法让 gorm 正确映射查询结果，应该定义明确的结构体
	// 或者使用 map 来接收动态结果
	var results []struct {
		Status string `json:"status"`
		Total  int64  `json:"total"`
	}
	// gorm 对大小写不敏感，甚至可以蛇形转驼峰（user_name => UserName）
	// 但 json 序列化/反序列化敏感，对不上时应使用标签 `json:"status"` 精确指定
	err := h.db.Model(&Task{}).Select("status, count(*) as total").Group("status").Scan(&results).Error
	if err != nil {
		// 修复：错误响应应该使用 gin.H 包装
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // 修复：出错后应该 return，避免继续执行
	}
	// ⚠️ 最佳实践：将查询结果转换为统一的格式，确保所有状态都有值（即使为0）
	stats := gin.H{
		"pending":   0,
		"running":   0,
		"completed": 0,
		"failed":    0,
	}
	for _, r := range results {
		stats[r.Status] = r.Total
	}
	c.JSON(http.StatusOK, stats)
}

// paramsToString 将 map 转换为 JSON 字符串（辅助函数）
func paramsToString(params map[string]interface{}) string {
	if params == nil {
		return "{}"
	}
	// ⚠️ 最佳实践：不要忽略 Marshal 的错误，虽然这里不太可能出错，但良好的习惯是检查错误
	data, err := json.Marshal(params)
	if err != nil {
		return "{}"
	}
	return string(data)
}
