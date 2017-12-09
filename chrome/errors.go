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
	var str string
	for _, errInfo := range e.errsInfo {
		str += fmt.Sprintf("target (%d): %s\n", errInfo.id, errInfo.err)
	}
	return str
}
