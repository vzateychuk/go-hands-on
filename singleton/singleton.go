package main

// Use the "sync" package for the Once API
import (
	"fmt"
	"sync"
)

// MyLogger is the struct we want to make a singleton
type MyLogger struct {
	loglevel int
}

// Log a message using the logger
func (l *MyLogger) Log(s string) {
	fmt.Println("Level -", l.loglevel, ":", s)
}

// SetLogLevel sets the log level of the logger
func (l *MyLogger) SetLogLevel(level int) {
	l.loglevel = level
}

var logger *MyLogger
var once sync.Once

// The getLoggerInstance function creates and provides global access to the Logger class instance
func getLoggerInstance() *MyLogger {
	// Use the sync.Once to enforce goroutine safety create singleton
	once.Do(func() {
		fmt.Println("Singleton : Creating logger instance")
		logger = &MyLogger{}
	})
	fmt.Println("Singleton : Returning existing instance")
	return logger
}
