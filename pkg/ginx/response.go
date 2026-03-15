package ginx

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewResponse 创建响应
func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewResponse(0, "success", data))
}

// Error 错误响应
func Error(c *gin.Context, err error, message ...string) {
	msg := "failed"
	if len(message) > 0 {
		msg = message[0]
	}
	if err != nil {
		msg += ": " + err.Error()
	}
	c.JSON(http.StatusInternalServerError, NewResponse(500, msg, nil))
}

// BadRequest 400 错误
func BadRequest(c *gin.Context, message string, data ...interface{}) {
	msg := message
	if len(data) > 0 {
		msg = formatMessage(message, data...)
	}
	c.JSON(http.StatusBadRequest, NewResponse(400, msg, nil))
}

// NotFound 404 错误
func NotFound(c *gin.Context, message string, data ...interface{}) {
	msg := message
	if len(data) > 0 {
		msg = formatMessage(message, data...)
	}
	c.JSON(http.StatusNotFound, NewResponse(404, msg, nil))
}

// formatMessage 格式化消息
func formatMessage(format string, a ...interface{}) string {
	if len(a) == 0 {
		return format
	}
	return fmt.Sprintf(format, a...)
}
