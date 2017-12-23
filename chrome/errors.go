package chrome

import (
	"fmt"
)

type errorInfo struct {
	err error
	id  int
}

type GoroutinesError struct {
	errsInfo []errorInfo
}

func (e *GoroutinesError) Error() string {
	str := "Goroutine error:\n"
	for _, errInfo := range e.errsInfo {
		str += fmt.Sprintf("instance (%d): %v\n", errInfo.id, errInfo.err)
	}
	return str
}
