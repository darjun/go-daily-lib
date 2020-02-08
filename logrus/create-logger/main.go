package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("info msg")
}
