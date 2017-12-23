package common

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func stackError(format string, a ...interface{}) error {
	skip := 3

	stackRows := strings.Split(string(debug.Stack()), "\n")
	stackRows = append(stackRows[:1], stackRows[1+skip*2:]...)

	stackStr := strings.Join(stackRows, "\n")

	return fmt.Errorf("%s\n%s", fmt.Sprintf(format, a...), stackStr)
}

func Error(msg string) error {
	return stackError(msg)
}

func Errorf(format string, a ...interface{}) error {
	return stackError(format, a...)
}
