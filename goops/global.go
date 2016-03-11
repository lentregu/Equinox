package goops

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var logger GoLogger

// C means context
type C map[string]string

func init() {

	logger = GoLogger{log: log.New(os.Stdout, "", 0)}
}

// Context serializes a map to json
func Context(context C) string {
	jsonMessage, _ := json.Marshal(context)
	return fmt.Sprintf("%s", jsonMessage)
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
