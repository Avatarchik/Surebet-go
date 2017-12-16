package db

import (
	"context"
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/chromedp/cdp"
	"github.com/korovkinand/surebetSearch/chrome"
	"github.com/korovkinand/surebetSearch/config/info"
	"github.com/korovkinand/surebetSearch/scanners"
	"github.com/korovkinand/surebetSearch/sites/posit"
	"log"
	"time"
)

type EventPairs = scanners.EventPairs

var dbPath = info.DbPath

func Collect() error {
	if err := info.Posit.Acc.Load(); err != nil {
		return err
	}

	positiveTargets := len(info.Posit.Acc.V)

	initLoads := make([]chromedp.Action, positiveTargets)
	for target := range initLoads {
		initLoads[target] = posit.InitLoad(info.Posit.Acc.V[target])
	}

	if err := chrome.InitPool(initLoads); err != nil {
		return err
	}
	defer chrome.ClosePool()

	var eventPairs EventPairs
	eventPairs.Load(dbPath)

	handleHtmls := make([]chromedp.Action, positiveTargets)
	for target := range handleHtmls {
		handleHtmls[target] = handleHtml(&eventPairs)
	}

	prevBackup := eventPairs.Length()
	workBegin := time.Now()
	for time.Since(workBegin) < info.Posit.Times.Limit {
		if err := chrome.RunActions(handleHtmls); err != nil {
			return err
		}

		if err := save(&eventPairs, &prevBackup); err != nil {
			return err
		}
		time.Sleep(info.Posit.Times.Sleep)
	}
	log.Print("Collecting ended")
	return nil
}

func handleHtml(eventPairs *EventPairs) chromedp.ActionFunc {
	var html string
	return func(ctx context.Context, c cdp.Handler) error {
		//Expects fixed code in fork of knq repo
		if err := chromedp.OuterHTML(info.Posit.Node, &html).Do(ctx, c); err != nil {
			return err
		}
		newPairs, err := posit.ParseHtml(&html)
		if err != nil {
			return err
		}
		eventPairs.AppendUnique(newPairs)

		return nil
	}
}

func save(eventPairs *EventPairs, prevBackup *int) error {
	curPairs := eventPairs.Length()
	log.Printf("Collected pairs: %d", curPairs)

	if curPairs-*prevBackup > 50 {
		if err := eventPairs.Save(dbPath + ".bkp"); err != nil {
			return err
		}
		*prevBackup = curPairs
		log.Println("Backup saved")
	}

	return eventPairs.Save(dbPath)
}
