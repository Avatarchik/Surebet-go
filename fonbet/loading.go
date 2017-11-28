package fonbet

import (
	"github.com/knq/chromedp"
	"github.com/knq/chromedp/cdp"
	"context"
	"log"
)

var url = "https://www.fonbet104.com/live/"
var expand = "#lineTableHeaderButton"
var expandAll = "#lineHeaderViewActionMenu > div:nth-child(6)"

func InitLoad() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady(expand),
		chromedp.Click(expand),
		chromedp.WaitReady(expandAll),
		chromedp.Click(expandAll),
		chromedp.WaitNotVisible(expandAll),
		chromedp.ActionFunc(func(ctxt context.Context, h cdp.Handler) error {
			log.Println("Fonbet loaded")
			return nil
		}),
	}
}
