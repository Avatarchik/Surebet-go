package main

import (
	"log"
	"surebetSearch/chrome"
	"github.com/knq/chromedp"
	"surebetSearch/common"
	"surebetSearch/bookmakers/fonbet"
)

func main() {
	address := "0.0.0.0"
	targetNumber := 1

	if err := chrome.RunPool(targetNumber, address); err != nil {
		log.Panic(err)
	}
	defer chrome.ClosePool()

	var initLoads []chromedp.Action
	for curTarget := 0; curTarget < targetNumber; curTarget++ {
		initLoads = append(initLoads, fonbet.InitLoad())
	}
	if errs := chrome.ExecActions(fonbet.LoadTimeout, initLoads); len(errs) != 0 {
		log.Panic(&chrome.GoroutineError{errs, fonbet.Url, "initLoad"})
	}
	var expandAll []chromedp.Action
	for curTarget := 0; curTarget < targetNumber; curTarget++ {
		expandAll = append(expandAll, fonbet.ExpandEvents())
	}
	if errs := chrome.ExecActions(fonbet.ExpandTimeout, expandAll); len(errs) != 0 {
		log.Panic(&chrome.GoroutineError{errs, fonbet.Url, "expandEvents"})
	}

	html := make([]string, targetNumber)

	var getHtmlAll []chromedp.Action
	for curTarget := 0; curTarget < targetNumber; curTarget++ {
		getHtmlAll = append(getHtmlAll, chrome.GetNodeHtml(fonbet.MainNode, &html[curTarget]))
	}
	if errs := chrome.ExecActions(fonbet.HtmlTimeout, getHtmlAll); len(errs) != 0 {
		log.Panic(&chrome.GoroutineError{errs, fonbet.Url, "GetHtml"})
	}

	common.SaveHtml(fonbet.Url, html[0])

	//if err := common.Benchmark(1, common.FuncsInfo{{}}); err != nil {
	//	log.Panic(err)
	//}
}
