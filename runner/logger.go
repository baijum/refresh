package runner

import (
	"fmt"
	"log"
	"os"
	"time"
)

type logFunc func(string, ...interface{})

var logger = log.New(os.Stderr, "", 0)

func newLogFunc(prefix string) func(string, ...interface{}) {
	color, clear := "", ""
	if settings["colors"] == "1" {
		color = fmt.Sprintf("\033[%sm", logColor(prefix))
		clear = fmt.Sprintf("\033[%sm", colors["reset"])
	}
	prefix = fmt.Sprintf("%-11s", prefix)

	return func(format string, v ...interface{}) {
		now := time.Now()
		timeString := fmt.Sprintf("%d:%d:%02d", now.Hour(), now.Minute(), now.Second())
		format = fmt.Sprintf("%s%s %s |%s %s", color, timeString, prefix, clear, format)
		logger.Printf(format, v...)
	}
}

type appLogWriter struct{}

func (a appLogWriter) Write(p []byte) (n int, err error) {
	appLog(string(p))

	return len(p), nil
}
