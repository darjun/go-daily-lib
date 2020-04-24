package main

import "go.uber.org/zap"

func main() {
	logger := zap.NewExample(zap.Fields(
		zap.Int("serverId", 90),
		zap.String("serverName", "awesome web"),
	))

	logger.Info("hello world")
}
