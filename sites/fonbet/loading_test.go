package fonbet

import (
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/surebetSearch/chrome"
	"testing"
)

func TestInitLoad(t *testing.T) {
	if err := chrome.RunPool(1); err != nil {
		t.Error(err)
	}
	defer chrome.ClosePool()

	if err := chrome.RunActions(InitLoad()); err != nil {
		t.Error(err)
	}

	var html string
	if err := chrome.RunActions(chromedp.OuterHTML(s.Node, &html)); err != nil {
		t.Error(err)
	}
}
