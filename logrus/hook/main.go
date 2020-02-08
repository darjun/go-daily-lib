package main

import (
	"github.com/sirupsen/logrus"
)

type AppHook struct {
	AppName string
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	return nil
}

func main() {
	h := &AppHook{AppName: "awesome-web"}
	logrus.AddHook(h)

	logrus.Info("info msg")
}
