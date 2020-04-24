package main

import "go.uber.org/zap"

func Output(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func main() {
	logger, _ := zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	Output("hello world")
}
