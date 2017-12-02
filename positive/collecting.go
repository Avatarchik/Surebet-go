package positive

import (
	"github.com/knq/chromedp"
	"surebetSearch/chrome"
	"log"
	"time"
	"os"
	"errors"
	"surebetSearch/common"
)

var filename = os.ExpandEnv("$GOPATH/src/surebetSearch/positive/collectedPairs")

func Collect() error {
	runReply, err := chrome.RunPool(1, "0.0.0.0")
	if err != nil {
		return err
	}
	defer chrome.ClosePool(runReply.Cancel, runReply.Pool)
	ctxt := runReply.Ctxt
	targets := runReply.Targets

	if errs := chrome.ExecActions(ctxt, targets, []chromedp.Action{
		InitLoad(),
	}); len(errs) != 0 {
		for _, err := range errs {
			return err
		}
	}

	var collectedPairs []EventPair
	common.LoadJson(filename, &collectedPairs)

	prevAmount := len(collectedPairs)

	for {
		if err := func() error {
			var html string
			if errs := chrome.ExecActions(ctxt, targets, []chromedp.Action{
				chrome.GetNodeHtml(MainNode, &html),
			}); len(errs) != 0 {
				for _, err := range errs {
					return err
				}
			}

			if len(html) == 0 {
				if errs := chrome.ExecActions(ctxt, targets, []chromedp.Action{
					chromedp.Tasks{chrome.SaveScn(loginUrl),
						chrome.WrapFunc(func() error {
							return common.SaveHtml(loginUrl, html)
						})},
				}); len(errs) != 0 {
					for _, err := range errs {
						return err
					}
				}
				return errors.New("got html with 0 length")
			}

			if err := ParseHtml(html, &collectedPairs); err != nil {
				return err
			}

			defer savePairs(&collectedPairs, &prevAmount)

			return nil
		}(); err != nil {
			return err
		}

		log.Println(len(collectedPairs))

		time.Sleep(2 * time.Second)
	}

	return nil
}

func savePairs(collectedPairs *[]EventPair, prevAmount *int) {
	if len(*collectedPairs) - *prevAmount > 100 {
		common.SaveJson(filename+".bkp", *collectedPairs)
		*prevAmount = len(*collectedPairs)
		log.Println("Backup saved")
	}

	common.SaveJson(filename, *collectedPairs)
}
