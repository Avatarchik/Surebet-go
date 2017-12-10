package positive

import (
	"github.com/korovkinand/chromedp"
	"log"
	"surebetSearch/chrome"
	"surebetSearch/common"
	"surebetSearch/config/paths"
)

var LoginUrl = "https://positivebet.com/en/user/login"

var loginSel = `#UserLogin_username`
var passSel = `#UserLogin_password`
var loginBtn = `#login-form > div.form-actions > button`

var liveBtn = `#yw0 > li:nth-child(2) > a`
var MainNode = `.grid-view > table`
var autoReloadBtn = `#formPanel > #btnAutoRefresh`
var changeAmountBar = `document.querySelector('#ddlPerPage').value = 30`

var Accounts []common.Account

func init() {
	if err := common.LoadJson(paths.PositiveAccounts, &Accounts); err != nil {
		log.Panic(err)
	}
}

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
