package positive

import (
	"github.com/jbowtie/gokogiri"
	"strings"
	"surebetSearch/dataBase/types"
)

func ParseHtml(html *string) ([]types.EventPair, error) {
	doc, err := gokogiri.ParseHtml([]byte(*html))
	if err != nil {
		return nil, err
	}
	defer doc.Free()

	trNodes, err := doc.Search(`//table/tbody/tr[not(@id = "")]`)
	if err != nil {
		return nil, err
	}

	var newPairs []types.EventPair

	for _, trNode := range trNodes {
		var eventPair types.EventPair

		bookmakers, err := trNode.Search(`.//td[3]/a`)
		if err != nil {
			return nil, err
		}

		for curMaker, bookmaker := range bookmakers {
			gotText, err := bookmaker.Search(`./text()`)
			if err != nil {
				return nil, err
			}

			if curMaker == 0 {
				eventPair.FirstEvent.BookMaker = gotText[0].String()
			} else {
				eventPair.SecondEvent.BookMaker = gotText[0].String()
			}
		}

		matchEvents, err := trNode.Search(`.//td[4]/a[@target="_blank"]`)
		if err != nil {
			return nil, err
		}
		for curMatch, matchEvent := range matchEvents {
			gotText, err := matchEvent.Search(`./text()`)
			if err != nil {
				return nil, err
			}

			teams := strings.Split(gotText[0].String(), " vs ")
			if curMatch == 0 {
				eventPair.FirstEvent.FirstTeam, eventPair.FirstEvent.SecondTeam = teams[0], teams[1]
			} else {
				eventPair.SecondEvent.FirstTeam, eventPair.SecondEvent.SecondTeam = teams[0], teams[1]
			}
		}

		newPairs = append(newPairs, eventPair)
	}
	return newPairs, nil
}
