package main

import (
	"fmt"
	
	"os"
	"time"
)

type Logger struct {
	file *os.File
}

func NewLogger(filename string) *Logger {
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return &Logger{file: file}
}

func (l *Logger) log(level string, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("[%s] [%s] %s", timestamp, level, msg)
	fmt.Println(logLine)
	fmt.Fprintln(l.file, logLine)
}

func (l *Logger) Info(msg string)  { l.log("INFO", msg) }
func (l *Logger) Warn(msg string)  { l.log("WARN", msg) }
func (l *Logger) Error(msg string) { l.log("ERROR", msg) }

func main() {
	logger := NewLogger("app.log")
	logger.Info("System started")
	logger.Warn("High memory usage")
	logger.Error("Pod crashed")
}
