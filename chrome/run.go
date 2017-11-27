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

	options := []chromedp.Option{chromedp.WithRunnerOptions(
		runner.Port(9222),
		runner.Flag("headless", true),
		runner.Flag("disable-gpu", true),
		runner.Flag("remote-debugging-address", address),
		runner.Flag("no-first-run", true),
		runner.Flag("no-default-browser-check", true),
		runner.UserDataDir(os.ExpandEnv("$HOME/ChromeDebug")),
		//runner.Flag("blink-settings", "imagesEnabled=false"),
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
