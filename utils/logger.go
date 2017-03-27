package utils

import (
    log "github.com/Sirupsen/logrus"
    "os"
)

var logger *log.Logger

func Logger() *log.Logger {
    if logger == nil {
        logger = log.New()
        logger.Level = log.DebugLevel
        logger.Out = os.Stderr
    }
    return logger
}