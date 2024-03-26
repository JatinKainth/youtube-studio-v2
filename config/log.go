package config

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var logInstance *logrus.Logger

func InitializeLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)

	logger.SetFormatter(&easy.Formatter{
		LogFormat:       "%time%[%lvl%][%rid%] %msg%\n",
		TimestampFormat: "2006-01-02T15:04:05.000",
	})

	logInstance = logger
	return logger
}

func GetLogger() *logrus.Logger {
	return logInstance
}
