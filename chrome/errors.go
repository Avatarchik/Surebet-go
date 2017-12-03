package chrome

import (
	"surebetSearch/common"
	"fmt"
	"log"
)

type GoroutineError struct {
	Errs []error
	Url  string
	Msg  string
}

func (e *GoroutineError) Error() string {
	for _, err := range e.Errs {
		log.Println(err)
	}
	siteName, err := common.GetSiteName(e.Url)
	if err != nil {
		siteName = "#can't parse url#"
	}
	return fmt.Sprintf("goroutine error: \"%s\" %s", siteName, e.Msg)
}
