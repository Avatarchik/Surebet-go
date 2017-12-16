package common

import (
	"strings"
	"time"
)

type Intervals struct {
	Sleep time.Duration
	Limit time.Duration
}

type SiteInfo struct {
	Acc   Accounts
	Times Intervals
	Url   string
	Node  string

	name    string
	prevUrl string
}

func (s *SiteInfo) Name() string {
	if s.prevUrl != "" && s.prevUrl == s.Url {
		return s.name
	}
	s.prevUrl = s.Url
	name, err := GetSiteName(s.Url)
	if err != nil {
		s.name = "InvalidName"
		return s.name
	}
	s.name = strings.Split(name, ".")[0]
	return s.name
}
