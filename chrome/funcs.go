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
		filename := siteName + ".png"
		if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
			return err
		}
		log.Printf("Screenshot saved to file %s", filename)
		return nil
	}
}

func SiteLoaded(s *common.SiteInfo, msg string) chromedp.ActionFunc {
	return func(ctx context.Context, c cdp.Handler) error {
		str := fmt.Sprintf("%s loaded", s.FullName())
		if msg != "" {
			str += fmt.Sprintf(": %s", msg)
		}
		log.Println(str)
		return nil
	}
}

func EventsLoaded(s *common.SiteInfo) chromedp.ActionFunc {
	return func(ctx context.Context, c cdp.Handler) error {
		log.Printf("%s events loaded", s.FullName())
		return nil
	}
}

func WrapFunc(fn func() error) chromedp.ActionFunc {
	return func(ctx context.Context, c cdp.Handler) error {
		return fn()
	}
}
