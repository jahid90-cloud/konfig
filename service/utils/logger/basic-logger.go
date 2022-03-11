package logger

import (
	"os"

	"github.com/fatih/color"
)

type basicLogger struct {
	level    LogLevel
	metadata string
}

func NewLogger() *basicLogger {
	return &basicLogger{
		// Default log level set to INFO
		level:    INFO,
		metadata: "",
	}
}

func (l *basicLogger) SetLogLevel(logLevel string) error {

	switch logLevel {
	case "DEBUG":
		l.level = DEBUG
	case "INFO":
		l.level = INFO
	case "WARN":
		l.level = WARN
	case "ERROR":
		l.level = ERROR
	case "FATAL":
		l.level = FATAL
	default:
		return ErrInvalidLogLevel
	}

	return nil
}

func (l *basicLogger) SetMetadata(metadata string) {
	l.metadata = metadata
}

func (l *basicLogger) Debug(args ...interface{}) error {
	if l.level <= DEBUG {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, ColorizerFunc(*color.New(color.FgBlue), "[debug]"))
		printArgs = append(printArgs, l.metadata)
		printArgs = append(printArgs, args...)

		return LoggerFunc(printArgs...)
	}

	return nil
}

func (l *basicLogger) Info(args ...interface{}) error {
	if l.level <= INFO {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, ColorizerFunc(*color.New(color.FgGreen), "[info]"))
		printArgs = append(printArgs, l.metadata)
		printArgs = append(printArgs, args...)

		return LoggerFunc(printArgs...)
	}

	return nil

}

func (l *basicLogger) Warn(args ...interface{}) error {
	if l.level <= WARN {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, ColorizerFunc(*color.New(color.FgYellow), "[warn]"))
		printArgs = append(printArgs, l.metadata)
		printArgs = append(printArgs, args...)

		return LoggerFunc(printArgs...)
	}

	return nil

}

func (l *basicLogger) Error(args ...interface{}) error {
	if l.level <= ERROR {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, ColorizerFunc(*color.New(color.FgRed), "[error]"))
		printArgs = append(printArgs, l.metadata)
		printArgs = append(printArgs, args...)

		return LoggerFunc(printArgs...)
	}

	return nil
}

func (l *basicLogger) Fatal(args ...interface{}) error {
	if l.level <= FATAL {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, ColorizerFunc(*color.New(color.FgHiRed), "[fatal]"))
		printArgs = append(printArgs, l.metadata)
		printArgs = append(printArgs, args...)

		LoggerFunc(printArgs...)
		os.Exit(1)
	}

	return nil
}

func (l *basicLogger) Debugf(format string, args ...interface{}) error {
	if l.level <= DEBUG {
		formatted := FormatterFunc(format, args...)
		return l.Debug(formatted)
	}

	return nil
}

func (l *basicLogger) Infof(format string, args ...interface{}) error {
	if l.level <= INFO {
		formatted := FormatterFunc(format, args...)
		return l.Info(formatted)
	}

	return nil
}

func (l *basicLogger) Warnf(format string, args ...interface{}) error {
	if l.level <= WARN {
		formatted := FormatterFunc(format, args...)
		return l.Warn(formatted)
	}

	return nil
}

func (l *basicLogger) Errorf(format string, args ...interface{}) error {
	if l.level <= ERROR {
		formatted := FormatterFunc(format, args...)
		return l.Error(formatted)
	}

	return nil
}

func (l *basicLogger) Fatalf(format string, args ...interface{}) error {
	if l.level <= FATAL {
		formatted := FormatterFunc(format, args...)
		return l.Fatal(formatted)
	}

	return nil
}
