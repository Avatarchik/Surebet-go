package main

import (
	"surebetSearch/chrome"
	"surebetSearch/fonbet"
	"github.com/knq/chromedp"
	"log"
)

func main() {
	address := "0.0.0.0"
	targetNumber := 5

	runReply, err := chrome.RunPool(targetNumber, address)
	if err != nil {
		log.Panic(err)
	}
	defer chrome.ClosePool(runReply.Cancel, runReply.Pool)
	ctxt := runReply.Ctxt
	targets := runReply.Targets

	if errs := chrome.ExecActions(ctxt, targets, []chromedp.Action{
		fonbet.InitLoad(),
		fonbet.InitLoad(),
		fonbet.InitLoad(),
		fonbet.InitLoad(),
		fonbet.InitLoad(),
	}); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		log.Panic("Goroutine error: InitLoad")
	}

	var html1, html2, html3, html4, html5 string
	if errs := chrome.ExecActions(ctxt, targets, []chromedp.Action{
		chrome.GetHtml(&html1),
		chrome.GetHtml(&html2),
		chrome.GetHtml(&html3),
		chrome.GetHtml(&html4),
		chrome.GetHtml(&html5),
	}); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		log.Panic("Goroutine error: GetHtml")
	}

	//if err := common.Benchmark(1, common.FuncsInfo{{}}); err != nil {
	//	log.Panic(err)
	//}
}
