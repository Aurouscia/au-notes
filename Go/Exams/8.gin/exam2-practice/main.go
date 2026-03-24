package main

import (
	"exam2-practice/middlewares"
	"exam2-practice/utils"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// 请在这里实现你的代码
// 提示：
// 1. 导入必要的包
// 2. 定义 User 结构体
// 3. 创建内存存储（map 或 slice）
// 4. 实现各个 Handler 函数
// 5. 设置路由并启动服务器

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"gte=0,lte=150"`
}

var idIncre = 0
var storage = make(map[int]*User, 10) // 线程不安全，无法并发
var mu = sync.Mutex{}

func main() {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(middlewares.MyLogger())

	g.GET("/users", getUsers)
	g.GET("/users/:id", getOneUser)
	g.POST("/users", createUser)
	g.PUT("/users/:id", updateUser)
	g.DELETE("/users/:id", deleteUser)

	g.Run()
}

func getUsers(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	// ❌ 该写法有问题：for-range 一个 map，第一返回值是 key，不能作为切片的 index 使用！
	// res := make([]User, len(storage)) // 注意第一个是 len 第二个是 cap
	// for idx, u := range storage {
	// 	res[idx] = *u
	// }

	res := make([]User, 0, len(storage))
	for _, u := range storage {
		res = append(res, *u)
	}

	utils.CreateResponseSuccess(ctx, res)
}

func getOneUser(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.CreateResponseFailed(ctx, http.StatusBadRequest, "user id invalid: %s", idStr)
		return
	}
	u := storage[id]
	if u == nil {
		utils.CreateResponseFailed(ctx, http.StatusNotFound, "user with id not found: %d", id)
		return
	}
	utils.CreateResponseSuccess(ctx, u)
}

func createUser(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	u := User{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		utils.CreateResponseFailed(ctx, http.StatusBadRequest, "user json invalid")
		return
	}
	u.ID = idIncre
	idIncre++
	storage[u.ID] = &u
	utils.CreateResponseSuccess(ctx, gin.H{})
}

func updateUser(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.CreateResponseFailed(ctx, http.StatusBadRequest, "user id invalid: %s", idStr)
		return
	}

	newU := User{}
	// ❌ 记得 Bind 要修改原 struct，必须传入指针而不是本体
	if err := ctx.ShouldBindJSON(&newU); err != nil {
		utils.CreateResponseFailed(ctx, http.StatusBadRequest, "user json invalid")
		return
	}

	u := storage[id]
	if u == nil {
		utils.CreateResponseFailed(ctx, http.StatusNotFound, "user with id not found: %d", id)
		return
	}
	u.Username = newU.Username
	u.Email = newU.Email
	u.Age = newU.Age
	utils.CreateResponseSuccess(ctx, gin.H{})
}

func deleteUser(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.CreateResponseFailed(ctx, http.StatusBadRequest, "user id invalid: %s", idStr)
		return
	}
	delete(storage, id)
	utils.CreateResponseSuccess(ctx, gin.H{})
}
