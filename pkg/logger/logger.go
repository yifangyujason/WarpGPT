package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	loglevel := os.Getenv("log_level")
	if loglevel == "" {
		loglevel = "info"
	}
	level, err := logrus.ParseLevel(loglevel)
	if err != nil {
		return
	}
	Log.SetLevel(level)

	Log.SetOutput(os.Stdout)

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
