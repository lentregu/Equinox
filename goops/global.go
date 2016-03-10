package goops

import (
	"log"
	"os"
)

var logger GoLogger

func init() {

	logger = GoLogger{log: log.New(os.Stdout, "", 0)}
}

// Info logs info messages
func Info(messages ...interface{}) {
	logger.info(messages...)
}

// Debug log debug
func Debug(messages ...interface{}) {
	logger.debug(messages...)
}

// Err log errors
func Err(messages ...interface{}) {
	logger.err(messages...)
}

// Fatal log errors
func Fatal(messages ...interface{}) {
	logger.fatal(messages...)
}
