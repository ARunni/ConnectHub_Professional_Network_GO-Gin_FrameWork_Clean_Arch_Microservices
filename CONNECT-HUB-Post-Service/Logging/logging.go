package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)



// InitLogrusLogger initializes a Logrus logger
func InitLogrusLogger(logFileName string) (*logrus.Logger, *os.File) {
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %s", err)
	}

	logger := logrus.New()
	logger.SetOutput(io.MultiWriter(os.Stdout, logFile))
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return logger, logFile
}
