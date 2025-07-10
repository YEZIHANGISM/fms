package log

import (
	"fms/configs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger  *zap.Logger        // 更快速的logger，强类型的结构化日志
	SLogger *zap.SugaredLogger // 更灵活的logger，支持任意类型的日志
)

func InitLogger() {
	level := parseLogLevel(configs.Config.Log.Level)
	zapConfig := zap.Config{
		Level:         zap.NewAtomicLevelAt(level),
		Encoding:      configs.Config.Log.Encoding,
		OutputPaths:   configs.Config.Log.OutputPaths,
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 设置时间格式为 ISO8601
	zapConfig.EncoderConfig.TimeKey = "time"                        // 设置时间键为 "time"

	Logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	SLogger = Logger.Sugar()
	defer Logger.Sync()  // 确保日志在程序退出前被写入
	defer SLogger.Sync() // 确保 SugaredLogger 的日志也被写入
}

func parseLogLevel(levelStr string) zapcore.Level {
	level := zapcore.InfoLevel // 默认级别为 InfoLevel
	if err := level.UnmarshalText([]byte(levelStr)); err != nil {
		// 如果解析失败，默认使用 InfoLevel
		level = zapcore.InfoLevel
	}
	return level
}
