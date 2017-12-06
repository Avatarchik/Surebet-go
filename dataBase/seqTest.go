package dataBase

import (
	"surebetSearch/dataBase/positive"
	"time"
	"log"
	"github.com/korovkinand/chromedp/runner"
	"github.com/korovkinand/chromedp"
	"context"
	"surebetSearch/dataBase/types"
	"os"
)

var options = []runner.CommandLineOption{
	runner.ExecPath("/usr/bin/google-chrome"),
	runner.Flag("headless", true),
	runner.Flag("disable-gpu", true),
	runner.Flag("no-first-run", true),
	runner.Flag("no-default-browser-check", true),
}

var html []string
var collectedPairs CollectedPairs

var filename = os.ExpandEnv("/home/ubuntu/go/src/surebetSearch/dataBase/collectedPairs")

func Test() {
	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create pool
	pool, err := chromedp.NewPool( /*chromedp.PoolLog(log.Printf, log.Printf, log.Printf)*/)
	if err != nil {
		log.Panic(err)
	}
	defer pool.Shutdown()

	targets := make([]*chromedp.Res, len(positive.Accounts))
	for i, account := range positive.Accounts {
		var err error
		targets[i], err = pool.Allocate(ctxt, options...)
		if err != nil {
			log.Panicf("url (%d) `%s` error: %v", i, account.Login, err)
		}
		defer targets[i].Release()

		if err := targets[i].Run(ctxt, positive.InitLoad(account)); err != nil {
			log.Panicf("url (%d) `%s` error: %v", i, account.Login, err)
		}
	}

	if err := collectedPairs.Load(filename); err != nil {
		log.Panic(err)
	}
	html = make([]string, len(positive.Accounts))

	for {
		for i, account := range positive.Accounts {
			collectPos(ctxt, i, account, targets[i])
		}
		time.Sleep(1 * time.Second)
	}
}

func collectPos(ctxt context.Context, id int, account types.Account, c *chromedp.Res) {
	if err := c.Run(ctxt, chromedp.Evaluate("document.documentElement.outerHTML", &html[id])); err != nil {
		log.Panicf("url (%d) `%s` error: %v", id, account.Login, err)
	}
	log.Printf("account: (%d) %s", id, account.Login)
	log.Printf("Html size: %d", len(html[id]))

	newPairs, err := positive.ParseHtml(&html[id])
	if err != nil {
		log.Panicf("url (%d) `%s` error: %v", id, account.Login, err)
	}

	collectedPairs.AppendUnique(newPairs)
	log.Printf("Collected pairs: %d", collectedPairs.Length())

	if err := collectedPairs.Save(filename); err != nil {
		log.Panic(err)
	}
}
