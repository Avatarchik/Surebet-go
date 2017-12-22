package config

import (
	"github.com/korovkinand/surebetSearch/common"
	"log"
)

var (
	fonbet = "fonbet"
	marat  = "marat"
	olimp  = "olimp"
	posit  = "posit"
)

var Sites = common.MakeSitesInfo()
var Accounts = common.MakeAccountsInfo()

func loadSites() error {
	if err := common.LoadJson(sitesConfigDir+fonbet, Sites.Fonbet); err != nil {
		return err
	}
	if err := common.LoadJson(sitesConfigDir+marat, Sites.Marat); err != nil {
		return err
	}
	if err := common.LoadJson(sitesConfigDir+olimp, Sites.Olimp); err != nil {
		return err
	}
	if err := common.LoadJson(sitesConfigDir+posit, Sites.Posit); err != nil {
		return err
	}
	return nil
}

func loadAccounts() error {
	var accounts []common.Account
	if err := common.LoadJson(accountsDir+posit, &accounts); err != nil {
		return err
	}
	Accounts.Posit.Set(accounts)
	Accounts.Posit.SetRange(Sites.Posit.Rng)

	return nil
}

func init() {
	if err := loadSites(); err != nil {
		log.Panic(err)
	}
	if err := loadAccounts(); err != nil {
		log.Panic(err)
	}
}
