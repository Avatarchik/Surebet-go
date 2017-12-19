package posit

import (
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/surebetSearch/chrome"
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/config"
)

var s *common.SiteInfo

func init() {
	s = config.Sites.Posit
}

func InitLoad(account common.Account) chromedp.Tasks {
	var res []byte
	return chromedp.Tasks{
		chromedp.Navigate(s.Url),
		chromedp.WaitVisible(s.Sel["login"]),
		chromedp.SendKeys(s.Sel["login"], account.Login),
		chromedp.SendKeys(s.Sel["pass"], account.Password),
		chromedp.Click(s.Sel["loginBtn"]),
		chromedp.WaitNotPresent(s.Sel["loginBtn"]),
		chromedp.Click(s.Sel["liveBtn"]),
		chromedp.WaitVisible(s.Node),
		chromedp.Click(s.Sel["autoReloadBtn"]),
		chromedp.Evaluate(s.Js["changeAmountBar"], &res),
		chrome.LogLoaded(s.FullName(), account.Login),
	}
}
