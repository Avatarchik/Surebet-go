package chrome

import (
	"github.com/knq/chromedp"
	"github.com/knq/chromedp/runner"
	"os"
	"context"
	"log"
)

type CDPInfo struct {
	Ctxt   context.Context
	C      *chromedp.CDP
	Cancel context.CancelFunc
}

func Run(address string, logEnabled bool) (*CDPInfo, error) {
	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	var options []chromedp.Option

	path := "/usr/bin/google-chrome"
	options = []chromedp.Option{chromedp.WithRunnerOptions(
		runner.Headless(path, 9222),
		runner.Flag("headless", true),
		runner.Flag("disable-gpu", true),
		runner.Flag("remote-debugging-address", address),
		//runner.Flag("blink-settings", "imagesEnabled=false"),
		runner.Flag("user-data-dir", os.ExpandEnv("$HOME/ChromeDebug")),
	)}

	if logEnabled {
		options = append(options, chromedp.WithLog(log.Printf))
	}
	// create chrome instance
	c, err := chromedp.New(ctxt, options...)
	if err != nil {
		return &CDPInfo{}, err
	}
	return &CDPInfo{ctxt, c, cancel}, nil
}

func Close(cdpInfo *CDPInfo) {
	// shutdown chrome
	if err := cdpInfo.C.Shutdown(cdpInfo.Ctxt); err != nil {
		log.Panic(err)
	}
	cdpInfo.Cancel()
}
