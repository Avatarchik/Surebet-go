package common

import (
	urlLib "net/url"
	"strings"
	"io/ioutil"
	"context"
	"github.com/knq/chromedp"
	"github.com/knq/chromedp/cdp"
	"log"
)

func GetSiteName(url string) (string, error) {
	u, err := urlLib.Parse(url)
	if err != nil {
		return "", err
	}
	return strings.Split(u.Hostname(), ".")[1], nil
}

func SaveScn(url string) chromedp.ActionFunc {
	return func(ctxt context.Context, c cdp.Handler) error {
		var buf []byte
		chromedp.CaptureScreenshot(&buf).Do(ctxt, c)
		siteName, err := GetSiteName(url)
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
