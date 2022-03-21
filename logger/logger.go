package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logger *logrus.Logger
}

func NewLogger() *Logger {
	loggrus := logrus.New()
	loggrus.SetFormatter(&logrus.JSONFormatter{})
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	loggrus.SetOutput(file)
	loggrus.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		Logger: loggrus,
	}
}

func (l *Logger) LogFatal(args interface{}) {
	l.Logger.Error(args)
	log.Println(args)
}

func (l *Logger) LogDebug(args interface{}) {
	l.Logger.Debug(args)
	log.Println(args)
}

func (l *Logger) LogInfo(args interface{}) {
	l.Logger.Info(args)
	log.Println(args)
}

func (l *Logger) LogWarn(args interface{}) {
	l.Logger.Warn(args)
	log.Println(args)
}
