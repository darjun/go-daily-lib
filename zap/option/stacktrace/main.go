package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func f1() {
	f2("hello world")
}

func f2(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}

func main() {
	logger, _ := zap.NewProduction(zap.AddStacktrace(zapcore.WarnLevel))
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	f1()
}
