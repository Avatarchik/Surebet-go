package fonbet

import (
	"github.com/knq/chromedp"
	"surebetSearch/common"
	"surebetSearch/chrome"
)

func Load(cdpInfo *chrome.CDPInfo) (string, error){
	ctxt, c := cdpInfo.Ctxt, cdpInfo.C
	url := "https://www.fonbet104.com/live/"
	expandBtn := "#lineTableHeaderButton"
	expandAll := "#lineHeaderViewActionMenu > div:nth-child(6)"
	var html string

	// run task list
	if err := c.Run(ctxt, clickButtons(url, expandBtn, expandAll, &html)); err != nil {
		return "", err
	}

	if err := common.SaveHtml(url, html); err != nil{
		return "", err
	}
	return html, nil
}

func clickButtons(url, btn, btnAll string, html *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(btn),
		chromedp.Click(btn, chromedp.NodeVisible),
		chromedp.WaitVisible(btnAll),
		chromedp.Click(btnAll, chromedp.NodeVisible),
		common.SaveScn(url),
		chromedp.OuterHTML("html", html),
	}
}