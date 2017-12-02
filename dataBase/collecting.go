package dataBase

import (
	"github.com/knq/chromedp"
	"log"
	"time"
	"os"
	"errors"
	"surebetSearch/chrome"
	"surebetSearch/common"
	"fmt"
	"surebetSearch/positive"
)

var filename = os.ExpandEnv("$GOPATH/src/surebetSearch/dataBase/collectedPairs")

func collect() error {
	positiveTargets := len(positive.Accounts)

	runReply, err := chrome.RunPool(positiveTargets, "0.0.0.0")
	if err != nil {
		return err
	}
	defer chrome.ClosePool(runReply.Cancel, runReply.Targets, runReply.Pool)
	ctxt := runReply.Ctxt
	targets := runReply.Targets

	var initLoads []chromedp.Action
	for curTarget := 0; curTarget < positiveTargets; curTarget++ {
		initLoads = append(initLoads, positive.InitLoad(curTarget))
	}

	if errs := chrome.ExecActions(ctxt, targets, initLoads); len(errs) != 0 {
		for _, err := range errs {
			return err
		}
	}

	var collectedPairs []common.EventPair
	common.LoadJson(filename, &collectedPairs)

	prevBackup := len(collectedPairs)
	for {
		if err := func() error {
			html := make([]string, positiveTargets)
			var getHtmlAll []chromedp.Action
			for curTarget := 0; curTarget < positiveTargets; curTarget++ {
				getHtmlAll = append(getHtmlAll, chrome.GetNodeHtml(positive.MainNode, &html[curTarget]))
			}

			if errs := chrome.ExecActions(ctxt, targets, getHtmlAll); len(errs) != 0 {
				for _, err := range errs {
					return err
				}
			}

			prevCollected := len(collectedPairs)
			for curTarget := 0; curTarget < positiveTargets; curTarget++ {
				if len(html[curTarget]) == 0 {
					urlName := fmt.Sprintf("https://positivebet%d.com", curTarget)
					if errs := chrome.ExecActions(ctxt, targets, []chromedp.Action{
						chromedp.Tasks{
							chrome.GetHtml(&html[curTarget]),
							chrome.SaveScn(urlName),
							chrome.WrapFunc(func() error {
								return common.SaveHtml(urlName, html[curTarget])
							})},
					}); len(errs) != 0 {
						for _, err := range errs {
							return err
						}
					}
					return errors.New("got html with 0 length")
				}

				if err := positive.ParseHtml(html[curTarget], &collectedPairs); err != nil {
					return err
				}
				log.Printf("Target# %d: %d", curTarget, len(collectedPairs)-prevCollected)
				prevCollected = len(collectedPairs)
			}

			defer savePairs(&collectedPairs, &prevBackup)

			return nil
		}(); err != nil {
			return err
		}

		log.Println(len(collectedPairs))

		time.Sleep(2 * time.Second)
	}

	return nil
}

func savePairs(collectedPairs *[]common.EventPair, prevBackup *int) {
	if len(*collectedPairs) - *prevBackup > 100 {
		common.SaveJson(filename+".bkp", *collectedPairs)
		*prevBackup = len(*collectedPairs)
		log.Println("Backup saved")
	}

	common.SaveJson(filename, *collectedPairs)
}

func Collect() {
	for {
		if err := collect(); err != nil {
			log.Println(err)
		}
		time.Sleep(5 * time.Second)
	}
}
