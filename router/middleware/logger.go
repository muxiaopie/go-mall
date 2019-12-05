package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/pkg/logger"
	"time"
)

var Log = logger.Logger

type Entry struct {
	Code 	  int
	Duration  string
	ClientIp  string
	Method    string
	Path      string
}

func Logger() gin.HandlerFunc {

	return func (c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		entry := Entry{
			Code:statusCode,
			Duration:fmt.Sprintf("%s",latency),
			ClientIp:clientIP,
			Method:method,
			Path:path,
		}
		if latency.Seconds() > 1.5 {
			Log.Warn(c.Request)
		}else{
			Log.Info(c.Request)
		}
		msg, _ := json.Marshal(entry)
		Log.Info(string(msg))
		/*Log.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)*/
	}
}