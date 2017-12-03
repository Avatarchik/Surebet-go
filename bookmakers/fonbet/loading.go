package fonbet

import (
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/chromedp/cdp"
	"context"
	"log"
	"time"
)

var LoadTimeout = 15 * time.Second
var ExpandTimeout = 5 * time.Second
var HtmlTimeout = 3 * time.Second

var Url = "https://www.fonbet104.com/live/"
var MainNode = `#lineTable > tbody`

var expand = "#lineTableHeaderButton"
var expandAll = "#lineHeaderViewActionMenu > .popupMenuItem:nth-child(6)"

var openNodesJs = `nodes = document.querySelectorAll('span[style="display: inline-block;"].detailArrowClose')
for (cur_node = 0; cur_node < nodes.length; cur_node++) {
    nodes[cur_node].click()
}`

func InitLoad() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(Url),
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

func ExpandEvents() chromedp.Tasks {
	var res []byte
	return chromedp.Tasks{
		chromedp.Evaluate(openNodesJs, &res),
	}
}

var expandAllJs = `document.querySelector('#lineTableHeaderButton').click();
document.querySelector('#lineHeaderViewActionMenu > div:nth-child(6)').click();`

func ExpandAllJs() chromedp.Tasks {
	var res []byte
	return chromedp.Tasks{
		chromedp.Evaluate(expandAllJs, &res),
	}
}

func ExpandAllCdp() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Click(expand),
		chromedp.WaitReady(expandAll),
		chromedp.Click(expandAll),
		chromedp.WaitNotVisible(expandAll),
	}
}
