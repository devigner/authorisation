package utils

import (
    log "github.com/Sirupsen/logrus"
    "os"
)



func CreateLogger() *log.Logger {
    logger := log.New()
    logger.Level = log.DebugLevel
    logger.Out = os.Stderr

    return logger
}