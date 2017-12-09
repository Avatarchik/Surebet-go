package dataBase

import (
	"github.com/korovkinand/chromedp"
	"log"
	"time"
	"surebetSearch/chrome"
	"surebetSearch/dataBase/positive"
	"surebetSearch/dataBase/types"
	"os"
	"github.com/korovkinand/chromedp/cdp"
	"context"
	"errors"
)

var filename = os.ExpandEnv("$GOPATH/src/surebetSearch/dataBase/collectedPairs")

func Collect() error {
	positiveTargets := len(positive.Accounts)

	if err := chrome.RunPool(positiveTargets); err != nil {
		return err
	}
	defer chrome.ClosePool()

	initLoads := make([]chromedp.Action, positiveTargets)
	for target := range initLoads {
		initLoads[target] = positive.InitLoad(positive.Accounts[target])
	}

	if err := chrome.RunActions(initLoads); err != nil {
		return err
	}

	var collectedPairs types.CollectedPairs
	collectedPairs.Load(filename)

	handleHtmls := make([]chromedp.Action, positiveTargets)
	for target := range handleHtmls {
		handleHtmls[target] = handleHtml(&collectedPairs)
	}

	prevBackup := collectedPairs.Length()
	for {
		if err := chrome.RunActions(handleHtmls); err != nil {
			return err
		}

		if err := save(&collectedPairs, &prevBackup); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}

	return errors.New("infinite loading ended")
}

func handleHtml(collectedPairs *types.CollectedPairs) chromedp.ActionFunc {
	var html string
	return func(ctx context.Context, c cdp.Handler) error {
		//Expects fixed code in fork of knq repo
		chromedp.OuterHTML(positive.MainNode, &html).Do(ctx, c)
		newPairs, err := positive.ParseHtml(&html)
		if err != nil {
			return err
		}
		collectedPairs.AppendUnique(newPairs)

		return nil
	}
}

func save(collectedPairs *types.CollectedPairs, prevBackup *int) error {
	curPairs := collectedPairs.Length()
	log.Printf("Collected pairs: %d", curPairs)

	if curPairs - *prevBackup > 50 {
		if err := collectedPairs.Save(filename + ".bkp"); err != nil {
			return err
		}
		*prevBackup = curPairs
		log.Println("Backup saved")
	}

	return collectedPairs.Save(filename)
}
