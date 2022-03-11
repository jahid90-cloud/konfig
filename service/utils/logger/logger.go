package logger

import (
	"errors"
)

var ErrInvalidLogLevel = errors.New("invalid log level")

type Logger interface {
	SetLogLevel(string) error
	SetMetadata(metadata string)
	Debug(...interface{}) error
	Debugf(string, ...interface{}) error
	Info(...interface{}) error
	Infof(string, ...interface{}) error
	Warn(...interface{}) error
	Warnf(string, ...interface{}) error
	Error(...interface{}) error
	Errorf(string, ...interface{}) error
	Fatal(...interface{}) error
	Fatalf(string, ...interface{}) error
}
