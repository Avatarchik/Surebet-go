package fonbet

import (
	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/common/types"
	"strconv"
	"strings"
)

type (
	bets    = types.Bets
	condBet = types.CondBet
	event   = types.Event
	events  = types.Events
)

type eventInfo struct {
	name      string
	isBlocked bool
}

var (
	handIDs  = []int{0, 1, 3}
	totalIDs = []int{0, 1, 2}
)

type rowInfo struct {
	class string
	node  xml.Node
}

const (
	trEvent        = "trEvent"
	trEventChild   = "trEventChild"
	trEventDetails = "trEventDetails"
)

func Parse(html string) (*types.Sports, error) {
	doc, err := gokogiri.ParseHtml([]byte(html))
	if err != nil {
		return nil, err
	}
	defer doc.Free()

	rowNodes, err := common.SearchAndCheck(doc, s.XPath["table"]+"/tr")
	if err != nil {
		return nil, err
	}

	var bookmaker types.Sports
	var allowedSports = map[string]*events{
		"1": &bookmaker.Soccer,
		"4": &bookmaker.Tennis,
		"2": &bookmaker.Hockey,
		"3": &bookmaker.Basket,
		"9": &bookmaker.Volley,
	}

	var rowsInfo []rowInfo
	var prevSport *events

	for _, rowNode := range rowNodes {
		rowClass := strings.Split(rowNode.Attr("class"), " ")
		if len(rowClass) == 1 {
			continue
		}
		if len(rowClass) != 3 {
			return nil, common.NewStackError("attribute @class in rowNode not found")
		}

		sportColorEnd := 10
		sportNum := rowClass[1][sportColorEnd:]
		sport, ok := allowedSports[sportNum]
		if !ok {
			continue
		}

		if rowClass[0] == trEvent {
			if err := appendEvent(rowsInfo, prevSport); err != nil {
				return nil, err
			}
			prevSport = sport
			rowsInfo = nil
		}
		rowsInfo = append(rowsInfo, rowInfo{rowClass[0], rowNode})
	}
	if err := appendEvent(rowsInfo, prevSport); err != nil {
		return nil, err
	}
	return &bookmaker, nil
}

func appendEvent(rowsInfo []rowInfo, sport *events) error {
	if rowsInfo != nil {
		event, err := parseEvent(rowsInfo)
		if err != nil {
			return err
		}
		sport.AppendNotEmpty(event)
	}
	return nil
}

func parseEvent(rowsInfo []rowInfo) (event, error) {
	node := rowsInfo[0].node

	evInfo, err := getEventInfo(node)
	if err != nil {
		return event{}, err
	}

	teams := strings.Split(evInfo.name, "—")
	if len(teams) != 2 || !strings.Contains(evInfo.name, "—") {
		return event{}, common.NewStackError("event name's struct has changed")
	}
	ev := event{Team1: teams[0], Team2: teams[1], Parts: []bets{}}

	partHandled := !evInfo.isBlocked
	if partHandled {
		bets, err := handleRow(node)
		if err != nil {
			return event{}, err
		}
		bets.Part = 0
		ev.Parts = append(ev.Parts, *bets)
	}

	for _, rowInfo := range rowsInfo[1:] {
		class, node := rowInfo.class, rowInfo.node
		switch class {
		case trEventDetails:
			if partHandled {
				if err := parseEventDetails(&ev, node); err != nil {
					return event{}, err
				}
			}
		case trEventChild:
			evInfo, err = getEventInfo(node)
			if err != nil {
				return event{}, err
			}

			partHandled = !evInfo.isBlocked && isPart(evInfo.name)
			if partHandled {
				strPartNum := string(evInfo.name[0])
				partNum, err := strconv.Atoi(strPartNum)
				if err != nil {
					return event{}, err
				}

				bets, err := handleRow(node)
				if err != nil {
					return event{}, err
				}
				bets.Part = partNum

				ev.Parts = append(ev.Parts, *bets)
			}
		}
	}
	return ev, nil
}

func parseEventDetails(event *event, node xml.Node) error {
	gridNodes, err := common.SearchAndCheck(node, "."+s.XPath["grid"])
	if err != nil {
		return err
	}

	bets := &(event.Parts)[len(event.Parts)-1]

	var allowedBets = map[string]*types.CondBets{
		"Hcap":          &bets.Hand,
		"Totals":        &bets.Total,
		"Team Totals-1": &bets.IndTotal1,
		"Team Totals-2": &bets.IndTotal2,
	}

	for _, gridNode := range gridNodes {
		captionNode, err := common.SearchAndCheck(gridNode, `.//thead/tr[1]/th/text()`)
		if err != nil {
			return err
		}

		caption := common.TrimSpaceNode(captionNode[0])
		condBets, ok := allowedBets[caption]
		if !ok {
			continue
		}

		gridRows, err := common.SearchAndCheck(gridNode, `.//tbody/tr`)
		if err != nil {
			return err
		}

		for _, gridRow := range gridRows {
			gridCols, err := common.SearchAndCheck(gridRow, `.//td`)
			if err != nil {
				return err
			}

			ids := totalIDs
			if caption == "Hcap" {
				ids = handIDs
			}
			condBet, err := handleCondBet(gridCols, ids)
			if err != nil {
				return err
			}
			condBets.AppendNotEmpty(condBet)
		}
	}
	return nil
}

func isPart(name string) bool {
	for _, partName := range []string{"half", "quarter", "set", "period"} {
		if strings.Contains(name, partName) {
			return true
		}
	}
	return false
}

func getEventInfo(rowNode xml.Node) (eventInfo, error) {
	eventTitle, err := common.SearchAndCheck(rowNode, "."+s.XPath["evName"])
	if err != nil {
		return eventInfo{}, err
	}

	eventNameNode, err := common.SearchAndCheck(eventTitle[0], `.//text()`)
	if err != nil {
		return eventInfo{}, err
	}

	evInfo := eventInfo{name: common.TrimSpaceNode(eventNameNode[1]), isBlocked: true}
	if eventTitle[0].Attr("class") != "eventBlocked" {
		evInfo.isBlocked = false
	}

	return evInfo, nil
}

func handleRow(rowNode xml.Node) (*bets, error) {
	var bets bets

	colNodes, err := common.SearchAndCheck(rowNode, `.//td`)
	if err != nil {
		return nil, err
	}

	var betAttrs = map[string]*types.Factor{
		"O1":  &bets.O1,
		"OX":  &bets.OX,
		"O2":  &bets.O2,
		"O1X": &bets.O1X,
		"O12": &bets.O12,
		"OX2": &bets.OX2,
	}
	for idx, betName := range []string{"O1", "OX", "O2", "O1X", "O12", "OX2"} {
		textNode, err := common.SearchText(colNodes[idx+3])
		if err != nil {
			return nil, err
		}

		if len(textNode) > 0 {
			factor, err := common.ParseFactor(textNode[0].String())
			if err != nil {
				return nil, err
			}
			*betAttrs[betName] = factor
		}
	}

	handBet, err := handleCondBet(colNodes[9:13], handIDs)
	if err != nil {
		return nil, err
	}

	bets.Hand.AppendNotEmpty(handBet)

	totalBet, err := handleCondBet(colNodes[13:16], totalIDs)
	if err != nil {
		return nil, err
	}
	bets.Total.AppendNotEmpty(totalBet)

	return &bets, nil
}

func handleCondBet(nodes []xml.Node, ids []int) (condBet, error) {
	var factors []types.Factor
	for _, id := range ids {
		res, err := common.SearchText(nodes[id])
		if err != nil {
			return condBet{}, err
		}
		length := len(res)
		if length > 1 {
			return condBet{}, common.NewStackError("structure has changed: cond bet")
		}
		if length == 1 {
			factor, err := common.ParseFactor(res[0].String())
			if err != nil {
				return condBet{}, err
			}
			factors = append(factors, factor)
		}
	}
	if len(factors) == 0 {
		return condBet{}, nil
	}
	if len(factors) != 3 {
		return condBet{}, common.NewStackError("structure has changed: cond bet")
	}
	return condBet{Cond: factors[0], V1: factors[1], V2: factors[2]}, nil
}
