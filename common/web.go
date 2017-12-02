package common

import (
	urlLib "net/url"
	"strings"
)

func GetSiteName(url string) (string, error) {
	u, err := urlLib.Parse(url)
	if err != nil {
		return "", err
	}
	parts := strings.Split(u.Hostname(), ".")
	siteName := parts[0]
	if parts[0] == "www" {
		siteName = parts[1]
	}
	return siteName, nil
}
