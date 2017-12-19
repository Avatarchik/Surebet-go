package posit

import (
	"github.com/jbowtie/gokogiri"
	"github.com/korovkinand/surebetSearch/common"
	"strings"
)

type EventPair = common.EventPair

func ParseHtml(html string) ([]EventPair, error) {
	doc, err := gokogiri.ParseHtml([]byte(html))
	if err != nil {
		return nil, err
	}
	defer doc.Free()

	trNodes, err := doc.Search(`//table/tbody/tr[not(@id = "")]`)
	if err != nil {
		return nil, err
	}

	var newPairs []EventPair

	for _, trNode := range trNodes {
		var eventPair EventPair

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
				eventPair.Event1.Site = gotText[0].String()
			} else {
				eventPair.Event2.Site = gotText[0].String()
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
				eventPair.Event1.Team1, eventPair.Event1.Team2 = teams[0], teams[1]
			} else {
				eventPair.Event2.Team1, eventPair.Event2.Team2 = teams[0], teams[1]
			}
		}

		newPairs = append(newPairs, eventPair)
	}
	return newPairs, nil
}
