package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 要求：
	// 1. 使用 SQLite 内存模式或文件模式
	// 2. 自动迁移 Task 模型
	// 3. 打印数据库初始化成功日志
	db, err := initDB()
	if err != nil {
		log.Fatal("数据库初始化失败:", err)
		return
	}
	log.Print("数据库初始化成功")

	// 要求：
	// 1. 创建 TaskScheduler，Worker 数量设置为 3
	// 2. 调度器会自动启动 Worker Goroutine
	scheduler := NewTaskScheduler(db, 3)

	// 要求：
	// 1. 创建 Gin 引擎
	// 2. 创建 TaskHandler
	// 3. 注册路由：
	//    - POST /api/tasks -> CreateTask
	//    - GET /api/tasks/:id -> GetTask
	//    - GET /api/tasks -> ListTasks
	//    - GET /api/tasks/stats -> GetTaskStats
	router := setupRouter(db, scheduler)

	// 启动 HTTP 服务器
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 优雅关闭处理
	go func() {
		// 等待中断信号
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("正在关闭服务器...")

		// 创建带超时的上下文
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// 关闭调度器
		if err := scheduler.Shutdown(ctx); err != nil {
			log.Println("调度器关闭失败:", err)
		}

		// 关闭 HTTP 服务器
		if err := srv.Shutdown(ctx); err != nil {
			log.Println("服务器关闭失败:", err)
		}
		log.Println("服务器已优雅关闭")
	}()

	log.Println("服务器启动在 :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("服务器启动失败:", err)
	}
}

// initDB 初始化数据库
func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"))
	if err != nil {
		db.AutoMigrate()
	}
	return db, err
}

// setupRouter 设置路由
func setupRouter(db *gorm.DB, scheduler *TaskScheduler) *gin.Engine {
	g := gin.Default()
	g.POST("/api/tasks", func(ctx *gin.Context) {
		task := Task{}
		ctx.BindJSON(&task)
		db.Create(&task)
		scheduler.SubmitTask(task.ID)
		ctx.JSON(http.StatusOK, task)
	})
	g.GET("/api/tasks/:id", func(ctx *gin.Context) {
		task := Task{}
		db.Find(&task, ctx.Param("id"))
		ctx.JSON(http.StatusOK, task)
	})
	g.GET("/api/tasks", func(ctx *gin.Context) {
		tasks := []Task{}
		db.Find(&tasks)
		ctx.JSON(http.StatusOK, tasks)
	})
	g.GET("/api/tasks/stats", func(ctx *gin.Context) {
		var results any
		db.Model(&Task{}).Select("ID, Status, count(*) as Total").Group("Status").Find(&results)
		ctx.JSON(http.StatusOK, results)
	})
	return nil
}
