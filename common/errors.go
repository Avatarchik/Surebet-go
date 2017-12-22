package common

import (
	"fmt"
	"runtime/debug"
)

type StackError struct {
	msg string
}

func (e *StackError) Error() string {
	return fmt.Sprintf("%s\n%v", e.msg, debug.Stack())
}

func NewStackError(msg string) *StackError {
	return &StackError{msg}
}

func NewStackErrorf(format string, a ...interface{}) *StackError {
	return NewStackError(fmt.Sprintf(format, a...))
}
