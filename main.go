package main

import (
	"log"
	"surebetSearch/common"
	"github.com/knq/chromedp"
	"surebetSearch/chrome"
)

func main() {
	address := "0.0.0.0"
	cdpInfo, err := chrome.Run(address, true,false)
	if err != nil{
		log.Panic(err)
	}
	defer chrome.Close(cdpInfo)

	//if _, err := fonbet.Load(cdpInfo); err != nil{
	//	log.Panic(err)
	//}

	ctxt, c := cdpInfo.Ctxt, cdpInfo.C
	url := "https://www.fonbet104.com/live/"
	expandBtn := "#lineTableHeaderButton"
	expandAll := "#lineHeaderViewActionMenu > .popupMenuItem:nth-child(6)"


	fonbetDocker := func() error{
		var html string
		if err := c.Run(ctxt, chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.WaitVisible(expandBtn),
			chromedp.Click(expandBtn),
			chromedp.WaitVisible(expandAll),
			chromedp.Click(expandAll),
			chromedp.WaitNotVisible(expandAll),
			//chromedp.WaitVisible(".trEventDetails", chromedp.ByQuery),
			chromedp.OuterHTML("html", &html)}); err != nil {
			return err
		}

		log.Printf("Html size: %d", len(html))

		return nil
	}

	common.Benchmark(3, []common.FuncInfo{{fonbetDocker, "fonbetDocker"}})
}
