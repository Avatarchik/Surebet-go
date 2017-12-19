package fonbet

import (
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/surebetSearch/chrome"
	"testing"
)

func TestInitLoad(t *testing.T) {
	if err := chrome.InitPool([]chromedp.Action{InitLoad()}); err != nil {
		t.Error(err)
	}
	defer chrome.ClosePool()

	var html string
	if err := chrome.RunActions([]chromedp.Action{chromedp.OuterHTML(s.Node, &html)}); err != nil {
		t.Error(err)
	}
}
