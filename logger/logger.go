package logger

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/YiZou89/bluebell/setting"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var sugaredLogger *zap.SugaredLogger
var logger *zap.Logger

func Init(conf *setting.LogConfig, mode string) (err error) {
	writeSyncer := getLogWriter(conf)
	encoder := getEncoder()

	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(conf.Level))
	if err != nil {
		return
	}

	var core zapcore.Core
	if mode == "dev" {
		// 开发模式，
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

		// 使用NewTee输出到文件以及终端
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)

	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	// AddCaller 添加调用方信息
	// AddCallerSkip(skip int)
	logger = zap.New(core, zap.AddCaller())

	//sugaredLogger = logger.Sugar()
	zap.ReplaceGlobals(logger)
	zap.L().Info("[logger] init success")
	return
}

// getLogWriter 使用lumberjack，添加日志切割归档功能
func getLogWriter(conf *setting.LogConfig) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize,    //最大文件大小 M
		MaxBackups: conf.MaxBackups, // 最大备份数量
		MaxAge:     conf.MaxAge,     //最大备份天数
		Compress:   false,           // 是否压缩
	}
	return zapcore.AddSync(lumberjackLogger)
}

// getEncoder 设置日志的编码格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 日志级别设置为大写，INFO，ERROR，WARN
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 日志时间格式修改为"2006-01-02T15:04:05.000Z0700"格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

// GinLogger 自定义日志，用于替换Gin框架中的log记录文件
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
