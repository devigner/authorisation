package utils

import (
    log "github.com/Sirupsen/logrus"
    "os"
)

var logInstance = createLogger()

func GetLogger() *log.Logger {
    return logInstance
}

func createLogger() *log.Logger {
    logger := log.New()
    logger.Level = log.DebugLevel
    logger.Out = os.Stderr

    return logger
}