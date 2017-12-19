package db

import (
	"fmt"
	"github.com/korovkinand/surebetSearch/common"
	"unicode"
)

func DelCyrillic(filename string) error {
	var evPairs []eventPair
	if err := common.LoadJson(filename, &evPairs); err != nil {
		return err
	}
	if err := common.SaveJson(filename+"Old", evPairs); err != nil {
		return err
	}

	var newPairs []eventPair
loop:
	for _, pair := range evPairs {
		teams := []string{pair.Event1.Team1, pair.Event1.Team2, pair.Event2.Team1, pair.Event2.Team2}
		for _, team := range teams {
			if isCyrillic(team) {
				continue loop
			}
		}
		newPairs = append(newPairs, pair)
	}

	fmt.Println(len(newPairs))

	return common.SaveJson(filename, newPairs)
}

func isCyrillic(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}
