package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	for i := 1; i <= 10; i++ {
		logrus.WithFields(logrus.Fields{
			"userid": i,
		}).Info("login")
		logrus.WithFields(logrus.Fields{
			"userid": i,
		}).Info("logoff")
	}
}
