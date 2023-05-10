package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLog(fileName string, maxSize int, maxBackups int, maxAge int, logLevel string) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	var l zapcore.Level = zapcore.InfoLevel

	switch logLevel {
	case "Debug":
		l = zapcore.DebugLevel
	case "Info":
		l = zapcore.InfoLevel
	case "Warn":
		l = zapcore.WarnLevel
	case "Error":
		l = zapcore.ErrorLevel
	case "DPanic":
		l = zapcore.DPanicLevel
	case "Panic":
		l = zapcore.PanicLevel
	case "Fatal":
		l = zapcore.FatalLevel
	default:
		l = zapcore.DebugLevel
	}

	core := zapcore.NewCore(encoder, writeSyncer, l)

	Logger = zap.New(core, zap.AddCaller())
}

func HexString(d []byte) string {
	var s string = ""
	for _, v := range d {
		t := fmt.Sprintf("0x%02x ", v)
		s = s + t
	}
	return s
}

func Debug(msg string, fields ...zapcore.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	Logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	Logger.Panic(msg, fields...)
}
