package config

import (
	"github.com/korovkinand/surebetSearch/common"
	"testing"
)

var sites *common.SitesInfo
var sitesList []*common.SiteInfo

func init() {
	sites = &Sites
	sitesList = []*common.SiteInfo{
		sites.Fonbet,
		sites.Marat,
		sites.Olimp,
		sites.Posit,
	}
}

func TestFullNames(t *testing.T) {
	fullNames := []string{
		"Fonbet",
		"Marathonbet",
		"Olimp",
		"Positivebet",
	}
	for siteNum, fullName := range fullNames {
		if fullName != sitesList[siteNum].FullName() {
			t.Fatal("full name isn't equal with known result")
		}
	}
}

func TestNames(t *testing.T) {
	names := []string{
		"fonbet",
		"marat",
		"olimp",
		"posit",
	}
	for siteNum, name := range names {
		if name != sitesList[siteNum].Name() {
			t.Fatal("name isn't equal with known result")
		}
	}
}
