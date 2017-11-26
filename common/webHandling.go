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
	return strings.Split(u.Hostname(), ".")[1], nil
}
