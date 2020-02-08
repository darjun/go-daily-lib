package main

import (
	"io/ioutil"

	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

func init() {
	hookConfig := logredis.HookConfig{
		Host:     "localhost",
		Key:      "mykey",
		Format:   "v0",
		App:      "aweosome",
		Port:     6379,
		Hostname: "localhost",
		DB:       0,
		TTL:      3600,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}
}

func main() {
	logrus.Info("just some info logging...")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"foo":    "bar",
		"this":   "that",
	}).Info("additional fields are being logged as well")

	logrus.SetOutput(ioutil.Discard)
	logrus.Info("This will only be sent to Redis")
}
