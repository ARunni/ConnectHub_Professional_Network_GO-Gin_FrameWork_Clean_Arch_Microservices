package helper

import (
	logging "github.com/ARunni/connectHub_gateway/Logging"

	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()
	return logrusLogger
}
