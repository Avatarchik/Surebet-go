package db

import (
	"context"
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/chromedp/cdp"
	"github.com/korovkinand/surebetSearch/chrome"
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/common/types"
	"github.com/korovkinand/surebetSearch/config"
	"github.com/korovkinand/surebetSearch/sites/posit"
	"log"
	"time"
)

type eventPair = common.EventPair
type set = types.SafeSet

var s *common.SiteInfo
var accounts *common.Accounts

func init() {
	s = config.Sites.Posit
	accounts = config.Accounts.Posit
}

func Collect(filename string) error {
	var rawPairs []eventPair
	if err := common.LoadJson(filename, &rawPairs); err != nil {
		return err
	}

	evPairs := types.NewSafeSet()
	evPairs.AddGen(rawPairs)

	positiveTargets := accounts.Size()

	initLoads := make([]chromedp.Action, positiveTargets)
	for target := range initLoads {
		initLoads[target] = posit.InitLoad(accounts.Values()[target])
	}

	if err := chrome.RunPool(len(initLoads)); err != nil {
		return err
	}
	defer chrome.ClosePool()

	if err := chrome.RunActions(initLoads...); err != nil {
		return err
	}

	handleHtmls := make([]chromedp.Action, positiveTargets)
	for target := range handleHtmls {
		handleHtmls[target] = handleHtml(evPairs)
	}

	prevBackup := evPairs.Size()
	workBegin := time.Now()
	for time.Since(workBegin) < s.Time.Limit {
		if err := chrome.RunActions(handleHtmls...); err != nil {
			return err
		}

		rawPairs := evPairs.ValuesGen(eventPair{}).([]eventPair)
		if err := save(filename, rawPairs, &prevBackup); err != nil {
			return err
		}
		time.Sleep(s.Time.Sleep)
	}
	log.Print("Collecting ended")
	return nil
}

func handleHtml(evPairs *set) chromedp.ActionFunc {
	var html string
	return func(ctx context.Context, c cdp.Handler) error {
		//Expects fixed code in fork of knq repo
		if err := chromedp.OuterHTML(s.Node, &html).Do(ctx, c); err != nil {
			return err
		}
		newPairs, err := posit.ParseHtml(html)
		if err != nil {
			return err
		}
		evPairs.AddGen(newPairs)

		return nil
	}
}

func save(filename string, evPairs []eventPair, prevBackup *int) error {
	amount := len(evPairs)
	log.Printf("Collected pairs: %d", amount)

	if amount-*prevBackup > 50 {
		if err := common.SaveJson(filename+".bkp", evPairs); err != nil {
			return err
		}
		*prevBackup = amount
		log.Println("Backup saved")
	}

	return common.SaveJson(filename, evPairs)
}
