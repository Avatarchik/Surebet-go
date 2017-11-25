package fonbet

import (
	"github.com/knq/chromedp"
	"surebetSearch/common"
	"surebetSearch/chrome"
)

func Load(cdpInfo *chrome.CDPInfo) (string, error) {
	ctxt, c := cdpInfo.Ctxt, cdpInfo.C
	url := "https://www.fonbet104.com/live/"
	expandBtn := "#lineTableHeaderButton"
	expandAll := "#lineHeaderViewActionMenu > .popupMenuItem:nth-child(6)"
	var html string

	// run task list
	if err := c.Run(ctxt, expandClick(url, expandBtn, expandAll, &html)); err != nil {
		return "", err
	}

	if err := common.SaveHtml(url, html); err != nil {
		return "", err
	}
	return html, nil
}

func expandClick(url, expandBtn, expandAll string, html *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady(expandBtn),
		chromedp.Click(expandBtn),
		chromedp.WaitReady(expandAll),
		chromedp.Click(expandAll),
		chromedp.WaitNotVisible(expandAll),
		common.SaveScn(url),
		chromedp.OuterHTML("html", html),
	}
}
