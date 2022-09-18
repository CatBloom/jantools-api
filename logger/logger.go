package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	log *zap.Logger
	err error
)

func LoggerInit() {
	log, err = zap.NewProduction()
	if err != nil {
		fmt.Println(err)
	}
}

func LoggerGet() *zap.Logger {
	return log
}

func LoggerClose() {
	log.Sync()
}

//middlewear
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.Info("Logger",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("elapsed", time.Since(start)),
		)
	}
}
