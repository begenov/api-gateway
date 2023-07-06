package logger

import (
	"runtime"
	"strconv"

	"github.com/begenov/api-gateway/internal/config"
	"github.com/begenov/api-gateway/pkg/errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Configuration provides configuration for zap logger

// Logger provides structure logging backed by uber zap logger
type Logger struct {
	*zap.Logger
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	c := caller.FullPath()
	details := runtime.FuncForPC(caller.PC)
	if details != nil {
		c = details.Name() + ":" + strconv.Itoa(caller.Line)
	}
	enc.AppendString(c)
}

// CreateLogger creates logrus logger
func CreateLogger(configuration config.LoggerConfig) *Logger {
	var logLevel zapcore.Level
	switch configuration.Level {
	case "DEBUG":
		logLevel = zapcore.DebugLevel
	case "INFO":
		logLevel = zapcore.InfoLevel
	case "WARN":
		logLevel = zapcore.WarnLevel
	case "ERROR":
		logLevel = zapcore.ErrorLevel
	default:
		panic("Invalid Log Level")
	}
	cfg := zap.Config{
		Encoding:         configuration.Format,
		Level:            zap.NewAtomicLevelAt(logLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: customCallerEncoder,

			StacktraceKey: "stacktrace",
		},
	}
	logger, _ := cfg.Build()
	return &Logger{logger}
}

// WithError return a new Logger with Error context
func (logger *Logger) WithError(err error) *Logger {
	if err == nil {
		return logger
	}

	var newLogger *zap.Logger
	if appError, ok := errors.IsAppError(err); ok {
		newLogger = logger.With(zap.String("error", appError.Error()),
			zap.String("errorstack", appError.ErrorStack()))
	} else {
		newLogger = logger.With(zap.String("error", err.Error()))
	}
	return &Logger{newLogger}
}
