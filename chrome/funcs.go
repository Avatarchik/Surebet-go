package chrome

import (
	"context"
	"fmt"
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/chromedp/cdp"
	"github.com/korovkinand/surebetSearch/common"
	"io/ioutil"
	"log"
)

func SaveScn(url string) chromedp.ActionFunc {
	return func(ctx context.Context, c cdp.Handler) error {
		var buf []byte
		if err := chromedp.CaptureScreenshot(&buf).Do(ctx, c); err != nil {
			return err
		}
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

func WaitClick(sel interface{}) chromedp.Action {
	return chromedp.Tasks{
		chromedp.WaitReady(sel),
		chromedp.Click(sel),
	}
}

func LogLoaded(site, msg string) chromedp.ActionFunc {
	return func(ctx context.Context, c cdp.Handler) error {
		str := fmt.Sprintf("%s loaded", site)
		if msg != "" {
			str += fmt.Sprintf(": %s", msg)
		}
		log.Println(str)
		return nil
	}
}

func WrapFunc(fn func() error) chromedp.ActionFunc {
	return func(ctx context.Context, c cdp.Handler) error {
		return fn()
	}
}
