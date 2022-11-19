package logger

import (
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 1 定义一下logger使用的常量
var (
	mode       string
	filename   string
	level      zapcore.Level
	maxSize    int
	maxAge     int
	maxBackups int
)

// 2 初始化Logger对象
func InitLogger() (err error) {
	// 读取配置
	mode = viper.GetString("log.mode")         //开发模式
	filename = viper.GetString("log.filename") // 日志存放路径
	//level    = viper.GetString("log.level")    // 日志级别
	level = zapcore.DebugLevel                   // 日志级别
	maxSize = viper.GetInt("log.max-size")       //最大存储大小
	maxAge = viper.GetInt("log.max-age")         //最大存储时间
	maxBackups = viper.GetInt("log.max-backups") //#备份数量

	// 创建Core三大件，进行初始化
	writeSyncer := getLogWriter(filename, maxSize, maxAge, maxBackups)
	encoder := getEncoder()
	// 创建核心-->如果是dev模式，就在控制台和文件都打印，否则就只写到文件中
	var core zapcore.Core
	if mode == "dev" {
		// 开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		// NewTee创建一个核心，将日志条目复制到两个或多个底层核心中。
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}

	//core := zapcore.NewCore(encoder, writeSyncer, level)
	// 创建 logger 对象
	log := zap.New(core, zap.AddCaller())
	// 替换全局的 logger, 后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(log)
	return
}

// 获取Encoder，给初始化logger使用的
func getEncoder() zapcore.Encoder {
	// 使用zap提供的 NewProductionEncoderConfig
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 时间的key
	encoderConfig.TimeKey = "time"
	// 级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 显示调用者信息
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 返回json 格式的 日志编辑器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 获取切割的问题，给初始化logger使用的
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	// 使用 lumberjack 归档切片日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 用于替换gin框架的Logger中间件，不传参数，直接这样写
func GinLogger(c *gin.Context) {
	logger := zap.L()
	start := time.Now()
	path := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	c.Next() // 执行视图函数
	// 视图函数执行完成，统计时间，记录日志
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

// GinRecovery 用于替换gin框架的Recovery中间件，因为传入参数，再包一层
func GinRecovery(stack bool) gin.HandlerFunc {
	logger := zap.L()
	return func(c *gin.Context) {
		defer func() {
			// defer 延迟调用，出了异常，处理并恢复异常，记录日志
			if err := recover(); err != nil {
				//  这个不必须，检查是否存在断开的连接(broken pipe或者connection reset by peer)---------开始--------
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				//httputil包预先准备好的DumpRequest方法
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// 如果连接已断开，我们无法向其写入状态
					c.Error(err.(error))
					c.Abort()
					return
				}
				//  这个不必须，检查是否存在断开的连接(broken pipe或者connection reset by peer)---------结束--------

				// 是否打印堆栈信息，使用的是debug.Stack()，传入false，在日志中就没有堆栈信息
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
				// 有错误，直接返回给前端错误，前端直接报错
				//c.AbortWithStatus(http.StatusInternalServerError)
				// 该方式前端不报错
				c.JSON(200, gin.H{"code": 500, "msg": "服务器异常"})
			}
		}()
		c.Next()
	}
}
