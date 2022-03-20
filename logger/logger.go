package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogrus() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}
