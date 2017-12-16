package db

import (
	"fmt"
	"github.com/korovkinand/surebetSearch/scanners"
	"unicode"
)

func delCyrillic(eventPairs EventPairs) error {
	var newPairs []scanners.EventPair
	for item := range eventPairs.Iter() {
		pair := item.V
		if !(checkCyrillic(pair.Event1.Team1) && checkCyrillic(pair.Event1.Team2) &&
			checkCyrillic(pair.Event2.Team1) && checkCyrillic(pair.Event2.Team2)) {
			newPairs = append(newPairs, pair)
		}
	}

	var newEventPairs EventPairs
	newEventPairs.AppendUnique(newPairs)

	fmt.Println(newEventPairs.Length())

	return newEventPairs.Save(dbPath + "new")
}

func checkCyrillic(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}
