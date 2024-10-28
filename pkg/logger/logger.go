package logger

import (
	"github.com/fatih/color"
)

type Logger struct {
	status  string
	message string
}

func NewLogger(status, message string) *Logger {
	return &Logger{
		status:  status,
		message: message,
	}
}

func (l *Logger) Log() {
	if l.status == "info" {
		color.Blue(l.message)
	}

	if l.status == "error" {
		color.Red(l.message)
	}

	if l.status == "success" {
		color.Green(l.message)
	}

	if l.status == "warning" {
		color.Yellow(l.message)
	}
}
