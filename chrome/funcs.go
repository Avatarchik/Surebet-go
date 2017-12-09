package chrome

import (
	"io/ioutil"
	"context"
	"log"
	"surebetSearch/common"
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/chromedp/cdp"
)

func SaveScn(url string) chromedp.ActionFunc {
	return func(ctx context.Context, c cdp.Handler) error {
		var buf []byte
		chromedp.CaptureScreenshot(&buf).Do(ctx, c)
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

func WrapFunc(fn func() error) chromedp.ActionFunc {
	return func(ctx context.Context, c cdp.Handler) error {
		return fn()
	}
}
