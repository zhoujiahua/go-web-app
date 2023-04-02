package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONResponse 将给定的数据、错误代码和消息编码为JSON，并将其发送到客户端。
func JSONResponse(c *gin.Context, code int, message string, data interface{}) {
	httpStatus := HTTPStatusFromCode(code)
	c.JSON(httpStatus, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// HTTPStatusFromCode 根据给定的错误代码返回适当的HTTP状态码。
func HTTPStatusFromCode(code int) int {
	switch code {
	case 0:
		return http.StatusOK
	case -1:
		return http.StatusInternalServerError
	case -2:
		return http.StatusBadRequest
	// 更多状态码...
	default:
		return http.StatusInternalServerError
	}
}
