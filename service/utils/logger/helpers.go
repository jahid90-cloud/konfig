package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var LoggerFunc = func(args ...interface{}) error {
	_, err := fmt.Fprintln(os.Stderr, args...)
	return err
}

var FormatterFunc = func(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

var ColorizerFunc = func(color color.Color, args ...interface{}) string {
	return color.Sprint(args...)
}
