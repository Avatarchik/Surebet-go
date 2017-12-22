package chrome

import (
	"fmt"
	"github.com/korovkinand/surebetSearch/common"
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
		str += fmt.Sprintf("instance (%d): %v\n", errInfo.id, errInfo.err)
	}
	return fmt.Sprint(common.NewStackError(str))
}
