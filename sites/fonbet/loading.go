package fonbet

import (
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/surebetSearch/chrome"
	"github.com/korovkinand/surebetSearch/config/info"
	"log"
)

var expand = "#lineTableHeaderButton"
var expandAll = "#lineHeaderViewActionMenu > .popupMenuItem:nth-child(6)"

var openNodesJs = `nodes = document.querySelectorAll('span[style="display: inline-block;"].detailArrowClose')
for (cur_node = 0; cur_node < nodes.length; cur_node++) {
    nodes[cur_node].click()
}`

func InitLoad() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(info.Fonbet.Url),
		chromedp.WaitReady(expand),
		chromedp.Click(expand),
		chromedp.WaitReady(expandAll),
		chromedp.Click(expandAll),
		chromedp.WaitNotVisible(expandAll),
		chrome.WrapFunc(func() error {
			log.Println("Fonbet loaded")
			return nil
		}),
	}
}

func ExpandEvents() chromedp.Action {
	var res []byte
	return chromedp.Evaluate(openNodesJs, &res)
}
