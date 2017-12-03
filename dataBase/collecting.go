package dataBase

import (
	"github.com/knq/chromedp"
	"log"
	"time"
	"os"
	"fmt"
	"surebetSearch/chrome"
	"surebetSearch/common"
	"surebetSearch/dataBase/positive"
	"surebetSearch/dataBase/types"
)

var filename = os.ExpandEnv("$GOPATH/src/surebetSearch/dataBase/collectedPairs")

func collect() error {
	positiveTargets := len(positive.Accounts)

	if err := chrome.RunPool(positiveTargets, "0.0.0.0"); err != nil {
		return err
	}
	defer chrome.ClosePool()

	var initLoads []chromedp.Action
	for curTarget := 0; curTarget < positiveTargets; curTarget++ {
		initLoads = append(initLoads, positive.InitLoad(curTarget))
	}

	if errs := chrome.ExecActions(positive.LoadTimeout, initLoads); len(errs) != 0 {
		return &chrome.GoroutineError{errs, positive.LoginUrl, "initLoad"}
	}

	var collectedPairs []types.EventPair
	common.LoadJson(filename, &collectedPairs)

	prevBackup := len(collectedPairs)
	for {
		html := make([]string, positiveTargets)
		var getHtmlAll []chromedp.Action
		for curTarget := 0; curTarget < positiveTargets; curTarget++ {
			getHtmlAll = append(getHtmlAll, chrome.GetNodeHtml(positive.MainNode, &html[curTarget]))
		}

		if errs := chrome.ExecActions(positive.HtmlTimeout, getHtmlAll); len(errs) != 0 {
			return &chrome.GoroutineError{errs, positive.LoginUrl, "getHtml"}
		}

		prevCollected := len(collectedPairs)
		for curTarget := 0; curTarget < positiveTargets; curTarget++ {
			if len(html[curTarget]) == 0 {
				log.Println("Got html with 0 length")
				url := fmt.Sprintf("https://positivebet%d.com", curTarget)
				if err := chrome.ReloadTarget(curTarget, positive.InitLoad(curTarget), url); err != nil {
					return err
				}
			}

			if err := positive.ParseHtml(html[curTarget], &collectedPairs); err != nil {
				return err
			}
			log.Printf("Target# %d: %d", curTarget, len(collectedPairs)-prevCollected)
			prevCollected = len(collectedPairs)
		}
		if err := savePairs(&collectedPairs, &prevBackup); err != nil {
			return err
		}

		log.Println(len(collectedPairs))
		time.Sleep(2 * time.Second)
	}
	return nil
}

func savePairs(collectedPairs *[]types.EventPair, prevBackup *int) error {
	if len(*collectedPairs) - *prevBackup > 100 {
		if err := common.SaveJson(filename+".bkp", *collectedPairs); err != nil {
			return err
		}
		*prevBackup = len(*collectedPairs)
		log.Println("Backup saved")
	}

	if err := common.SaveJson(filename, *collectedPairs); err != nil {
		return err
	}
	return nil
}

func Collect() {
	for {
		if err := collect(); err != nil {
			log.Println(err)
		}
		time.Sleep(5 * time.Second)
	}
}
