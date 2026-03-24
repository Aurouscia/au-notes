package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateResponseSuccess(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

func CreateResponseFailed(ctx *gin.Context, code int, format string, params ...any) {
	ctx.JSON(code, gin.H{
		"code":    0,
		"message": fmt.Sprintf(format, params...),
	})
}
