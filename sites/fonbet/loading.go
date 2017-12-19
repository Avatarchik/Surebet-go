package fonbet

import (
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/surebetSearch/chrome"
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/config"
)

var s *common.SiteInfo

func init() {
	s = config.Sites.Fonbet
}

func InitLoad() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(s.Url),
		chrome.WaitClick(s.Sel["expand"]),
		chrome.WaitClick(s.Sel["expandAll"]),
		chromedp.WaitNotVisible(s.Sel["expandAll"]),
		chrome.LogLoaded(s.FullName(), ""),
	}
}

func ExpandEvents() chromedp.Action {
	var res []byte
	return chromedp.Evaluate(s.Js["openNodes"], &res)
}
