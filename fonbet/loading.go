package fonbet

import (
	"github.com/knq/chromedp"
	"surebetSearch/chrome"
	"github.com/knq/chromedp/client"
)

var url = "https://www.fonbet104.com/live/"
var expand = "#lineTableHeaderButton"
var expandAll = "#lineHeaderViewActionMenu > div:nth-child(6)"

func LoadClient(cdpInfo *chrome.CDPInfo, targetClient *client.Client) error {
	target, err := targetClient.NewPageTarget(cdpInfo.Ctxt)
	if err != nil {
		return err
	}
	cdpInfo.C.AddTarget(cdpInfo.Ctxt, target)
	cdpInfo.C.SetHandlerByID(target.GetID())

	if err := cdpInfo.C.Run(cdpInfo.Ctxt, initLoad(url, expand, expandAll)); err != nil {
		return err
	}
	return nil
}

func Load(cdpInfo *chrome.CDPInfo) error {
	var targetID string
	if err := cdpInfo.C.Run(cdpInfo.Ctxt, cdpInfo.C.NewTarget(&targetID)); err != nil {
		return err
	}
	cdpInfo.C.SetHandlerByID(targetID)

	// run task list
	if err := cdpInfo.C.Run(cdpInfo.Ctxt, initLoad(url, expand, expandAll)); err != nil {
		return err
	}
	return nil
}

func initLoad(url, expandBtn, expandAll string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady(expandBtn),
		chromedp.Click(expandBtn),
		chromedp.WaitReady(expandAll),
		chromedp.Click(expandAll),
		chromedp.WaitNotVisible(expandAll),
	}
}
