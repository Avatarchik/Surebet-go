package common

import (
	"fmt"
	"github.com/jbowtie/gokogiri/xml"
	"strconv"
	"strings"
)

func ReformatSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func ParseFactor(s string) (Factor, error) {
	value, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return Factor(0), err
	}
	return Factor(value), err
}

func SearchAndCheck(node xml.Node, xpath string) ([]xml.Node, error) {
	res, err := node.Search(xpath)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("node not found, xpath: %s", xpath)
	}

	return res, nil
}

func SearchText(node xml.Node) ([]xml.Node, error) {
	return node.Search(`.//text()`)
}

func TrimSpaceNode(node xml.Node) string {
	return strings.TrimSpace(node.String())
}
