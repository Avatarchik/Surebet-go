package dataBase

import (
	"github.com/korovkinand/chromedp"
	"log"
	"time"
	"surebetSearch/chrome"
	"surebetSearch/common"
	"surebetSearch/dataBase/positive"
	"surebetSearch/dataBase/types"
	"errors"
)

func Collect() error {
	positiveTargets := len(positive.Accounts)

	if err := chrome.RunPool(positiveTargets, "0.0.0.0"); err != nil {
		return err
	}
	defer chrome.ClosePool()

	var initLoads []chromedp.Action
	for _, account := range positive.Accounts {
		initLoads = append(initLoads, positive.InitLoad(account))
	}

	if errs := chrome.ExecActions(positive.LoadTimeout, initLoads); len(errs) != 0 {
		return &chrome.GoroutineError{errs, positive.LoginUrl, "initLoad"}
	}

	var collectedPairs []types.EventPair
	common.LoadJson(filename, &collectedPairs)

	//prevBackup := len(collectedPairs)
	for {
		html := make([]string, positiveTargets)
		var getHtmlAll []chromedp.Action
		for curTarget := 0; curTarget < positiveTargets; curTarget++ {
			getHtmlAll = append(getHtmlAll, chrome.GetNodeHtml(positive.MainNode, &html[curTarget]))
		}

		if errs := chrome.ExecActions(positive.HtmlTimeout, getHtmlAll); len(errs) != 0 {
			return &chrome.GoroutineError{errs, positive.LoginUrl, "getHtml"}
		}

		for curTarget := 0; curTarget < positiveTargets; curTarget++ {
			if len(html[curTarget]) == 0 {
				return errors.New("got html with 0 length")
			}

			//newPairs, err := positive.ParseHtml(&html[curTarget])
			//if err != nil {
			//	return err
			//}
			//
			//uniques := UniqueAndDiff(&collectedPairs, newPairs)
			//if err := savePairs(collectedPairs, &prevBackup); err != nil {
			//	return err
			//}
			//log.Printf("Target# %d: %d", curTarget, uniques)
		}
		log.Println(len(collectedPairs))
		time.Sleep(2 * time.Second)
	}
	return nil
}

func savePairs(collectedPairs []types.EventPair, prevBackup *int) error {
	pairsAmount := len(collectedPairs)
	if pairsAmount - *prevBackup > 50 {
		if err := common.SaveJson(filename+".bkp", collectedPairs); err != nil {
			return err
		}
		*prevBackup = pairsAmount
		log.Println("Backup saved")
		// Temp code
		if err := common.SaveJson(filename, collectedPairs); err != nil {
			return err
		}
	}
	return nil
}
