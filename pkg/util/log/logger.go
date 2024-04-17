package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var LogrusObj *logrus.Logger

// 这里的日志图方便，直接全部在终端输出了，一般来说都会在文件中输出
func InitLog() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: true,
	})
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	/*
		TODO:
			加个hook形成ELK体系
			到时候再加
	*/
	LogrusObj = logger
}
