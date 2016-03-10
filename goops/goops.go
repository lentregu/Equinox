package goops

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// GoLogger type defines the info to be written in a log trace
type GoLogger struct {
	// Level is the log level
	Level string `json:"lvl"`
	// time is the timestamp at the log has been written
	Time string `json:"time"`
	// Msg is the log message
	Msg interface{} `json:"msg"`

	log *log.Logger
}

const (
	debug = "DEBUG"
	info  = "INFO"
	warn  = "WARN"
	err   = "ERROR"
	fatal = "FATAL"
)

// New creates a logger
func New() GoLogger {

	return GoLogger{log: log.New(os.Stdout, "", 0)}

}

// Info logs info messages
func (l GoLogger) Info(messages ...interface{}) {
	fmtString := ""
	var message interface{}
	var i int
	args := make([]interface{}, len(messages)-1)
	for i, message = range messages {
		if value, ok := message.(string); !ok {
			fmtString += "%v,"
			args[i] = message
		} else {
			fmtString += value + "\n"
			args = append(args[:i], messages[i+1:]...)
			break
		}
	}

	fmt.Printf(fmtString, args...)

}

// Debug log debug
func (l GoLogger) Debug(message interface{}) {
	l.Level = debug
	l.jsonLog(message)
}

// Error log errors
func (l GoLogger) Error(message interface{}) {
	l.Level = err
	l.jsonLog(message)
}

// Error log errors
func (l GoLogger) Fatal(message interface{}) {
	l.Level = err
	l.jsonLogFatal(message)
}

func (l GoLogger) jsonLog(message interface{}) {
	l.Time = time.Now().Format(time.RFC3339)
	l.Msg = message
	jsonMessage, _ := json.Marshal(l)
	l.log.Printf("%s", jsonMessage)
}

func (l GoLogger) jsonLogFatal(message interface{}) {
	l.Time = time.Now().Format(time.RFC3339)
	l.Msg = message
	jsonMessage, _ := json.Marshal(l)
	l.log.Fatal("%s", jsonMessage)
}
