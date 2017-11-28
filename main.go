package main

import (
	"surebetSearch/chrome"
	"surebetSearch/fonbet"
	"github.com/knq/chromedp"
	"log"
)

func main() {
	address := "0.0.0.0"
	targetNumber := 10

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

	html := make([]string, targetNumber)
	if errs := chrome.ExecActions(ctxt, targets, []chromedp.Action{
		chrome.GetHtml(&html[0]),
		chrome.GetHtml(&html[1]),
		chrome.GetHtml(&html[2]),
		chrome.GetHtml(&html[3]),
		chrome.GetHtml(&html[4]),
		chrome.GetHtml(&html[5]),
		chrome.GetHtml(&html[6]),
		chrome.GetHtml(&html[7]),
		chrome.GetHtml(&html[8]),
		chrome.GetHtml(&html[9]),
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
