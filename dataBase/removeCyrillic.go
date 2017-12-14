package dataBase

import (
	"fmt"
	"surebetSearch/config/paths"
	"surebetSearch/dataBase/types"
	"unicode"
)

func RemoveCyrillic(collectedPairs types.CollectedPairs) error {
	var newPairs []types.EventPair
	for item := range collectedPairs.Iter() {
		pair := item.V
		if !(checkCyrillic(pair.FirstEvent.FirstTeam) && checkCyrillic(pair.FirstEvent.SecondTeam) &&
			checkCyrillic(pair.SecondEvent.FirstTeam) && checkCyrillic(pair.SecondEvent.SecondTeam)) {
			newPairs = append(newPairs, pair)
		}
	}

	var newCollectedPairs types.CollectedPairs
	newCollectedPairs.AppendUnique(newPairs)

	fmt.Println(newCollectedPairs.Length())

	return newCollectedPairs.Save(paths.CollectedPairs + "new")
}

func checkCyrillic(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}
