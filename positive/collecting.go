package positive

import (
	"github.com/knq/chromedp"
	"log"
	"time"
	"os"
	"errors"
	"surebetSearch/chrome"
	"surebetSearch/common"
	"fmt"
)

var filename = os.ExpandEnv("$GOPATH/src/surebetSearch/positive/collectedPairs")

func Collect() error {
	targetNumber := 6

	runReply, err := chrome.RunPool(targetNumber, "0.0.0.0")
	if err != nil {
		return err
	}
	defer chrome.ClosePool(runReply.Cancel, runReply.Pool)
	ctxt := runReply.Ctxt
	targets := runReply.Targets

	var initLoads []chromedp.Action
	for curTarget := 0; curTarget < targetNumber; curTarget++ {
		initLoads = append(initLoads, InitLoad(curTarget))
	}

	if errs := chrome.ExecActions(ctxt, targets, initLoads); len(errs) != 0 {
		for _, err := range errs {
			return err
		}
	}

	var collectedPairs []EventPair
	common.LoadJson(filename, &collectedPairs)

	prevBackup := len(collectedPairs)
	for {
		if err := func() error {
			html := make([]string, targetNumber)
			var getHtmlAll []chromedp.Action
			for curTarget := 0; curTarget < targetNumber; curTarget++ {
				getHtmlAll = append(getHtmlAll, chrome.GetNodeHtml(MainNode, &html[curTarget]))
			}

			if errs := chrome.ExecActions(ctxt, targets, getHtmlAll); len(errs) != 0 {
				for _, err := range errs {
					return err
				}
			}

			prevCollected := len(collectedPairs)
			for curTarget := 0; curTarget < targetNumber; curTarget++ {
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

				if err := ParseHtml(html[curTarget], &collectedPairs); err != nil {
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

func savePairs(collectedPairs *[]EventPair, prevBackup *int) {
	if len(*collectedPairs) - *prevBackup > 100 {
		common.SaveJson(filename+".bkp", *collectedPairs)
		*prevBackup = len(*collectedPairs)
		log.Println("Backup saved")
	}

	common.SaveJson(filename, *collectedPairs)
}