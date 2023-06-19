package logging

import (
	logger "github.com/sirupsen/logrus"
)

func PanicWithError(msg string, err error) {
	logger.Panic(msg, err)
}
