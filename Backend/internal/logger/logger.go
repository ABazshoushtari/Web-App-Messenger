package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func Init() {
	logLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.Lock(os.Stdout), logLevel)
	logger = zap.New(core).Sugar()
}

func Logger() *zap.SugaredLogger {
	return logger
}
