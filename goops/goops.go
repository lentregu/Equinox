package goops

import (
	"encoding/json"
	"fmt"
	"log"
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

// Info logs info messages
func (l GoLogger) parse(messages ...interface{}) string {
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

	return fmt.Sprintf(fmtString, args...)

}

// Info logs info messages
func (l GoLogger) info(messages ...interface{}) {
	l.Level = info
	l.jsonLog(l.parse(messages...))

}

// Debug log debug
func (l GoLogger) debug(messages ...interface{}) {
	l.Level = debug
	l.jsonLog(l.parse(messages...))
}

// Error log errors
func (l GoLogger) err(messages ...interface{}) {
	l.Level = err
	l.jsonLog(l.parse(messages...))
}

// Fatal log errors
func (l GoLogger) fatal(messages ...interface{}) {
	l.Level = err
	l.jsonLogFatal(l.parse(messages...))
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
	l.log.Fatalf("%s", jsonMessage)
}
