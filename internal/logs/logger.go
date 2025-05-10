package logs

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	error  *log.Logger
	writer io.Writer
}

func NewLogger(pfx string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, pfx, log.Ldate|log.Ltime)

	return &Logger{
		debug:  log.New(writer, "DEBUG", logger.Flags()),
		info:   log.New(writer, "INFO", logger.Flags()),
		warn:   log.New(writer, "WARN", logger.Flags()),
		error:  log.New(writer, "ERROR", logger.Flags()),
		writer: writer,
	}
}

// Debug Create Non-formated Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

// Info Create Non-formated Logs
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

// Warn Create Non-formated Logs
func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

// Error Create Non-formated Logs
func (l *Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}

// Debugf Create Format Enabled Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

// Infof Create Format Enabled Logs
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

// Warnf Create Format Enabled Logs
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

// Errorf Create Format Enabled Logs
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}
