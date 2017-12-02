package positive

import (
	"github.com/knq/chromedp"
	"github.com/knq/chromedp/cdp"
	"context"
	"log"
)

var loginUrl = "https://positivebet.com/en/user/login"

var login = []string{
	"volosha123@gmail.com",
	"kokozhina@gmail.com",
	"marshytv@ya.ru",
	"ilya00@gmail.com",
	"kolyan312@gmail.com",
	"petya146@gmail.com",
}
var pass = "1q1w1e1r"

var loginSel = `#UserLogin_username`
var passSel = `#UserLogin_password`
var loginBtn = `#login-form > div.form-actions > button`

var liveBtn = `#yw0 > li:nth-child(2) > a`
var MainNode = `.grid-view > table`
var autoReloadBtn = `#formPanel > #btnAutoRefresh`
var changeAmountBar = `document.querySelector('#ddlPerPage').value = 30`

func InitLoad(targetNumber int) chromedp.Tasks {
	var res []byte
	return chromedp.Tasks{
		chromedp.Navigate(loginUrl),
		chromedp.WaitVisible(loginSel),
		chromedp.SendKeys(loginSel, login[targetNumber]),
		chromedp.SendKeys(passSel, pass),
		chromedp.Click(loginBtn),
		chromedp.WaitNotPresent(loginBtn),
		chromedp.Click(liveBtn),
		chromedp.WaitVisible(MainNode),
		chromedp.Click(autoReloadBtn),
		chromedp.Evaluate(changeAmountBar, &res),
		chromedp.ActionFunc(func(ctxt context.Context, h cdp.Handler) error {
			log.Printf("Loaded positive# %d", targetNumber)
			return nil
		}),
	}
}
