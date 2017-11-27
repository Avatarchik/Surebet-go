package fonbet

import (
	"github.com/knq/chromedp"
	"surebetSearch/chrome"
	"sync"
	"context"
	"log"
	"github.com/knq/chromedp/runner"
)

var url = "https://www.fonbet104.com/live/"
var expand = "#lineTableHeaderButton"
var expandAll = "#lineHeaderViewActionMenu > div:nth-child(6)"

func LoadPool(ctxt context.Context, wg *sync.WaitGroup, pool *chromedp.Pool, id int) {
	defer wg.Done()

	options := []runner.CommandLineOption{
		runner.ExecPath("/usr/bin/google-chrome"),
		runner.Flag("headless", true),
		runner.Flag("disable-gpu", true),
		runner.Flag("remote-debugging-address", "0.0.0.0"),
		runner.Flag("no-first-run", true),
		runner.Flag("no-default-browser-check", true),
	}

	// allocate
	c, err := pool.Allocate(ctxt, options...)
	if err != nil {
		log.Printf("id# %d error: %v", id, err)
		return
	}
	defer c.Release()

	if err := c.Run(ctxt, initLoad(url, expand, expandAll)); err != nil {
		log.Printf("id# %d error: %v", id, err)
		return
	}
	var html string
	if err := c.Run(ctxt, chrome.GetHtml(&html)); err != nil {
		log.Printf("id# %d error: %v", id, err)
		return
	}
	log.Printf("Html size: %d", len(html))
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
