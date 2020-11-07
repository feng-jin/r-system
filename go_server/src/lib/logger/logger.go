package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var logger = log.New()

type LogConfig struct {
	Level string `toml:"level"`
	Dir   string `toml:"dir"`
}

func InitLogger(logConfig *LogConfig) {
	fmt.Println(*logConfig)
	f, _ := os.OpenFile(logConfig.Dir, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	logger.SetOutput(f)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	switch logConfig.Level {
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}
