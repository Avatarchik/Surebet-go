package fonbet

import (
	"github.com/korovkinand/surebetSearch/sites"
	"testing"
)

func TestLoad(t *testing.T) {
	if err := sites.TestLoadEvents(s, InitLoad(), LoadEvents()); err != nil {
		t.Fatal(err)
	}
}
