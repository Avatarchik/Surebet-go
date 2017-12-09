package positive

import (
	"github.com/korovkinand/chromedp"
	"log"
	"surebetSearch/chrome"
	"surebetSearch/common"
)

var LoginUrl = "https://positivebet.com/en/user/login"

var Accounts = []common.Account{
	{"volosha123@gmail.com", "1q1w1e1r"},
	{"kokozhina@gmail.com", "1q1w1e1r"},
	{"marshytv@ya.ru", "1q1w1e1r"},
	{"ilya00@gmail.com", "1q1w1e1r"},
	{"kolyan312@gmail.com", "1q1w1e1r"},
	{"petya146@gmail.com", "1q1w1e1r"},
	{"lester0578@gmail.com", "1q1w1e1r"},
}

var loginSel = `#UserLogin_username`
var passSel = `#UserLogin_password`
var loginBtn = `#login-form > div.form-actions > button`

var liveBtn = `#yw0 > li:nth-child(2) > a`
var MainNode = `.grid-view > table`
var autoReloadBtn = `#formPanel > #btnAutoRefresh`
var changeAmountBar = `document.querySelector('#ddlPerPage').value = 30`

func InitLoad(account common.Account) chromedp.Tasks {
	var res []byte
	return chromedp.Tasks{
		chromedp.Navigate(LoginUrl),
		chromedp.WaitVisible(loginSel),
		chromedp.SendKeys(loginSel, account.Login),
		chromedp.SendKeys(passSel, account.Password),
		chromedp.Click(loginBtn),
		chromedp.WaitNotPresent(loginBtn),
		chromedp.Click(liveBtn),
		chromedp.WaitVisible(MainNode),
		chromedp.Click(autoReloadBtn),
		chromedp.Evaluate(changeAmountBar, &res),
		chrome.WrapFunc(func() error {
			log.Printf("Positive loaded: %s", account.Login)
			return nil
		}),
	}
}
