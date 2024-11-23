package logger

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

// AccessLog 请求日志
type AccessLog struct {
	Duration     string
	Method       string
	Url          string
	status       int
	RequestBody  string
	ResponseBody string
}

// MiddleWareBuilder 日志打印中间件
type MiddleWareBuilder struct {
	allowRequestBody  bool
	allowResponseBody bool
	loggerFunc        func(ctx context.Context, log *AccessLog)
}

// NewBuilder 创建日志打印中间件构造器
func NewBuilder(fn func(ctx context.Context, log *AccessLog)) *MiddleWareBuilder {
	return &MiddleWareBuilder{
		loggerFunc: fn,
	}
}

// AllowRequestBody 允许打印请求体
func (b *MiddleWareBuilder) AllowRequestBody() *MiddleWareBuilder {
	b.allowRequestBody = true
	return b
}

// AllowResponseBody 允许打印响应体
func (b *MiddleWareBuilder) AllowResponseBody() *MiddleWareBuilder {
	b.allowResponseBody = true
	return b
}

// responseWriter 装饰gin.ResponseWriter 获取响应体
type responseWriter struct {
	gin.ResponseWriter
	log *AccessLog
}

func (b *MiddleWareBuilder) Build() gin.HandlerFunc {
	start := time.Now()
	return func(ctx *gin.Context) {
		url := ctx.Request.URL.String()
		if len(url) > 1024 {
			url = url[:1024] + "..."
		}
		log := AccessLog{
			Url:    url,
			Method: ctx.Request.Method,
		}
		if b.allowRequestBody && ctx.Request.Body != nil {
			body, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
			if len(body) > 1024 {
				body = body[:1024]
			}
			log.RequestBody = string(body)
		}
		if b.allowResponseBody {
			ctx.Writer = &responseWriter{
				ResponseWriter: ctx.Writer,
				log:            &log,
			}
		}
		log.Duration = time.Since(start).String()
		defer func() {
			b.loggerFunc(ctx, &log)
		}()
		ctx.Next()
	}
}

// Write 获取响应体
func (r *responseWriter) Write(data []byte) (int, error) {
	r.log.ResponseBody = string(data)
	return r.ResponseWriter.Write(data)
}

// WriteString 获取响应体
func (r *responseWriter) WriteString(data string) (int, error) {
	r.log.RequestBody = data
	return r.ResponseWriter.WriteString(data)
}

// WriteHeader 获取响应状态码
func (r *responseWriter) WriteHeader(statuscode int) {
	r.log.status = statuscode
	r.ResponseWriter.WriteHeader(statuscode)
}
