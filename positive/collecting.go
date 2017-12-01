package positive

import (
	"github.com/knq/chromedp"
	"surebetSearch/chrome"
	"io/ioutil"
	"encoding/json"
	"log"
	"time"
	"os"
)

var filename = os.ExpandEnv("$GOPATH/src/surebetSearch/positive/collectedPairs.json")


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

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var collectedPairs []EventPair
	err = json.Unmarshal(data, &collectedPairs)
	if err != nil {
		return err
	}

	for {
		if err := func() error{
			var html string
			if errs := chrome.ExecActions(ctxt, targets, []chromedp.Action{
				chrome.GetNodeHtml(MainNode, &html),
			}); len(errs) != 0 {
				for _, err := range errs {
					return err
				}
			}

			if err := ParseHtml(html, &collectedPairs); err != nil {
				return err
			}

			data, err = json.Marshal(collectedPairs)
			if err != nil {
				return err
			}
			defer saveJson(data)

			return nil
		}(); err != nil {
			return err
		}

		log.Println(len(collectedPairs))

		time.Sleep(2 * time.Second)
	}


	return nil
}

func saveJson(data []byte) {
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		log.Panic(err)
	}
}