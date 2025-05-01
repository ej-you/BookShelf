package logger

import (
	"log"
	"os"
)

var _ Logger = (*logger)(nil)

type Logger interface {
	Debug(args ...any)
	Debugf(template string, args ...any)
	Print(args ...any)
	Printf(template string, args ...any)
	Error(args ...any)
	Errorf(template string, args ...any)
}

// Logger implementation
type logger struct {
	debugLog *log.Logger
	infoLog  *log.Logger
	errorLog *log.Logger
}

// Logger constructor
func NewLogger() Logger {
	return &logger{
		debugLog: log.New(os.Stdout, "[DEBUG]\t", log.Ldate|log.Ltime),
		infoLog:  log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime),
	}
}

func (l logger) Debug(args ...any) {
	l.debugLog.Print(args...)
}

func (l logger) Debugf(template string, args ...any) {
	l.debugLog.Printf(template, args...)
}

func (l logger) Print(args ...any) {
	l.infoLog.Println(args...)
}

func (l logger) Printf(template string, args ...any) {
	l.infoLog.Printf(template, args...)
}

func (l logger) Error(args ...any) {
	l.errorLog.Println(args...)
}

func (l logger) Errorf(template string, args ...any) {
	l.errorLog.Printf(template, args...)
}
