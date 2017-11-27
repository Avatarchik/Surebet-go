package chrome

import (
	"io/ioutil"
	"context"
	"log"
	"surebetSearch/common"
	"github.com/knq/chromedp"
	"github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/client"
)

var targetClient = client.New()

func SaveScn(url string) chromedp.ActionFunc {
	return func(ctxt context.Context, c cdp.Handler) error {
		var buf []byte
		chromedp.CaptureScreenshot(&buf).Do(ctxt, c)
		siteName, err := common.GetSiteName(url)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(siteName+".png", buf, 0644); err != nil {
			return err
		}
		log.Println("Screenshot saved")
		return nil
	}
}

func GetHtml(html *string) chromedp.ActionFunc {
	return func(ctxt context.Context, c cdp.Handler) error {
		chromedp.OuterHTML("html", html).Do(ctxt, c)
		return nil
	}
}

func PrintTargets(cdpInfo *CDPInfo) {
	targets, err := targetClient.ListTargets(cdpInfo.Ctxt)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Opened tabs: %d", len(targets)-1)
	log.Println("CLient: ")
	for _, target := range targets {
		log.Println(target)
	}
}
