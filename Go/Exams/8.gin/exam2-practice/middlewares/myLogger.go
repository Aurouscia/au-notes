package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := ctx.Request.URL
		method := ctx.Request.Method
		t := time.Now()

		ctx.Next()

		// code := ctx.Request.Response.StatusCode // ❌ 本次响应的状态码不是这么获取的，这个是客户端请求时携带的响应（比如重定向时的上一个响应）
		code := ctx.Writer.Status()
		if code == 0 {
			code = 200
		}
		tSpan := time.Since(t)
		tSpanMs := tSpan.Milliseconds()
		fmt.Printf("HTTP %s %s returned %d in %dms\n", method, url, code, tSpanMs)
	}
}
