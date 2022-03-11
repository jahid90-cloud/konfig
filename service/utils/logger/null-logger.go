package logger

// nullLogger A null logger implementation; all ops are no-op
type nullLogger struct{}

func NewNullLogger() *nullLogger {
	return &nullLogger{}
}

func (l *nullLogger) SetLogLevel(logLevel string) error {
	return nil
}

func (l *nullLogger) SetMetadata(metadata string) {}

func (l *nullLogger) Debug(args ...interface{}) error {
	return nil
}

func (l *nullLogger) Info(args ...interface{}) error {
	return nil
}

func (l *nullLogger) Warn(args ...interface{}) error {
	return nil
}

func (l *nullLogger) Error(args ...interface{}) error {
	return nil
}

func (l *nullLogger) Fatal(args ...interface{}) error {
	return nil
}

func (l *nullLogger) Debugf(format string, args ...interface{}) error {
	return nil
}

func (l *nullLogger) Infof(format string, args ...interface{}) error {
	return nil
}

func (l *nullLogger) Warnf(format string, args ...interface{}) error {
	return nil
}

func (l *nullLogger) Errorf(format string, args ...interface{}) error {
	return nil
}

func (l *nullLogger) Fatalf(format string, args ...interface{}) error {
	return nil
}
