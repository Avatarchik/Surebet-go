package positive

import (
	"github.com/jbowtie/gokogiri"
	"strings"
	"surebetSearch/dataBase/types"
)

func ParseHtml(html string, collectedPairs *[]types.EventPair) error {
	doc, err := gokogiri.ParseHtml([]byte(html))
	if err != nil {
		return err
	}

	trNodes, err := doc.Search(`//table/tbody/tr[not(@id = "")]`)
	if err != nil {
		return err
	}

	for _, trNode := range trNodes {
		var eventPair types.EventPair

		bookmakers, err := trNode.Search(`.//td[3]/a`)
		if err != nil {
			return err
		}

		for curMaker, bookmaker := range bookmakers {
			gotText, err := bookmaker.Search(`./text()`)
			if err != nil {
				return err
			}

			if curMaker == 0 {
				eventPair.FirstEvent.BookMaker = gotText[0].String()
			} else {
				eventPair.SecondEvent.BookMaker = gotText[0].String()
			}
		}

		matchEvents, err := trNode.Search(`.//td[4]/a[@target="_blank"]`)
		if err != nil {
			return err
		}
		for curMatch, matchEvent := range matchEvents {
			gotText, err := matchEvent.Search(`./text()`)
			if err != nil {
				return err
			}

			teams := strings.Split(gotText[0].String(), " vs ")
			if curMatch == 0 {
				eventPair.FirstEvent.FirstTeam, eventPair.FirstEvent.SecondTeam = teams[0], teams[1]
			} else {
				eventPair.SecondEvent.FirstTeam, eventPair.SecondEvent.SecondTeam = teams[0], teams[1]
			}
		}

		*collectedPairs = append(*collectedPairs, eventPair)
	}

	*collectedPairs = uniq(*collectedPairs)

	return nil
}

func uniq(list []types.EventPair) []types.EventPair {
	unique_set := make(map[types.EventPair]bool, len(list))
	for _, x := range list {
		unique_set[x] = true
	}
	result := make([]types.EventPair, 0, len(unique_set))
	for x := range unique_set {
		result = append(result, x)
	}
	return result
}
