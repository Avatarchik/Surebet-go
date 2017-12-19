package common

import (
	"time"
	"unicode"
)

type Time struct {
	Sleep time.Duration `json:"sleep,omitempty"`
	Limit time.Duration `json:"limit,omitempty"`
}

type Map map[string]string

type SiteInfo struct {
	Url   string `json:"url"`
	Time  *Time  `json:"time,omitempty"`
	Rng   *Range `json:"range,omitempty"`
	Node  string `json:"node,omitempty"`
	Sel   Map    `json:"sel,omitempty"`
	Js    Map    `json:"js,omitempty"`
	XPath Map    `json:"xpath,omitempty"`

	name, fullName string
}

type SitesInfo struct {
	Fonbet,
	Marat,
	Olimp,
	Posit *SiteInfo
}

func MakeSitesInfo() SitesInfo {
	return SitesInfo{
		&SiteInfo{},
		&SiteInfo{},
		&SiteInfo{},
		&SiteInfo{},
	}
}

type AccountsInfo struct {
	Fonbet,
	Marat,
	Olimp,
	Posit *Accounts
}

func MakeAccountsInfo() AccountsInfo {
	return AccountsInfo{
		&Accounts{},
		&Accounts{},
		&Accounts{},
		&Accounts{},
	}
}

func isVowel(r rune) bool {
	switch unicode.ToLower(r) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func (s *SiteInfo) FullName() string {
	if s.fullName == "" {
		name, err := GetSiteName(s.Url)
		if err != nil {
			name = "InvalidUrl"
		}
		s.fullName = name
	}
	return s.fullName
}

func (s *SiteInfo) Name() string {
	if s.name == "" {
		name := s.FullName()

		const maxLen = 6
		if len(name) > maxLen {
			end := maxLen - 1
			if isVowel([]rune(name)[end-1]) {
				end++
			}
			name = name[:end]
		}
		s.name = name
	}
	return s.name
}