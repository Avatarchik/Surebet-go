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
		chromedp.Click(s.Sel["expand"]),
		chromedp.Click(s.Sel["expandAll"]),
		chromedp.WaitNotVisible(s.Sel["expandAll"]),
		chrome.SiteLoaded(s, ""),
	}
}

func LoadEvents() chromedp.Tasks {
	var res []byte
	return chromedp.Tasks{
		chromedp.Evaluate(s.Js["openNodes"], &res),
		chrome.EventsLoaded(s),
	}
}
