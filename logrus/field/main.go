package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age": 18,
	}).Info("info msg")
}