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
		// ⚠️最佳实践：log.Fatal 会调用 os.Exit(1) 终止程序，后面的 return 永远不会执行
		// 应该去掉多余的 return 语句
		log.Fatal("数据库初始化失败:", err)
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
		// ⚠️ 最佳实践：使用 signal.NotifyContext 可以更方便地处理信号和取消
		// 使用 signal.Notify 将操作系统信号注册到指定的 channel，当程序接收到指定的信号时，信号会被发送到该 channel。
		// SIGINT 中断信号
		// SIGTERM 终止信号
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("正在关闭服务器...")

		// 创建带超时的上下文
		// 注意：这里使用同一个 ctx 先关闭调度器再关闭服务器，如果调度器关闭耗时较长，
		// 可能会导致服务器关闭时 ctx 已经超时。建议分开处理或使用更长的超时时间。
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// 关闭调度器
		if err := scheduler.Shutdown(ctx); err != nil {
			log.Println("调度器关闭失败:", err)
		}

		// 关闭 HTTP 服务器
		// Shutdown 会优雅地关闭服务器，等待现有连接处理完成
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
	db, err := gorm.Open(sqlite.Open("tasks.db"))
	if err != nil {
		return nil, err
	}
	// ❌ 检查 AutoMigrate 的错误返回值
	if err := db.AutoMigrate(&Task{}); err != nil {
		return nil, err
	}
	return db, nil
}

// setupRouter 设置路由
func setupRouter(db *gorm.DB, scheduler *TaskScheduler) *gin.Engine {
	g := gin.Default()

	handler := NewTaskHandler(db, scheduler)

	g.POST("/api/tasks", handler.CreateTask)
	g.GET("/api/tasks/:id", handler.GetTask)
	g.GET("/api/tasks", handler.ListTasks)
	g.GET("/api/tasks/stats", handler.GetTaskStats)
	return g
}
