package positive


import (
	"github.com/jbowtie/gokogiri"
	"strings"
)

type MatchEvent struct {
	BookMaker string `json:"BookMaker"`
	FirstTeam string `json:"FirstTeam"`
	SecondTeam string `json:"SecondTeam"`
}

type EventPair struct {
	FirstEvent MatchEvent `json:"FirstEvent"`
	SecondEvent MatchEvent `json:"SecondEvent"`
}


func ParseHtml(html string, collectedPairs *[]EventPair) error {
	doc, err := gokogiri.ParseHtml([]byte(html))
	if err != nil {
		return err
	}

	trNodes, err := doc.Search(`//table/tbody/tr[not(@id = "")]`)
	if err != nil {
		return err
	}

	for _, trNode := range trNodes {
		var eventPair EventPair

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

func uniq(list []EventPair) []EventPair {
	unique_set := make(map[EventPair]bool, len(list))
	for _, x := range list {
		unique_set[x] = true
	}
	result := make([]EventPair, 0, len(unique_set))
	for x := range unique_set {
		result = append(result, x)
	}
	return result
}