package fonbet

import (
	"fmt"
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/common/types"
	"github.com/korovkinand/surebetSearch/config"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestKnownResult(t *testing.T) {
	data, err := ioutil.ReadFile(config.FonbetSamplesDir + "KnownResult.html")
	if err != nil {
		t.Fatal(err)
	}
	html := string(data)
	bookmaker, err := Parse(html)
	if err != nil {
		t.Fatal(err)
	}

	sports := []*events{
		&bookmaker.Soccer,
		&bookmaker.Tennis,
		&bookmaker.Hockey,
		&bookmaker.Basket,
		&bookmaker.Volley,
	}
	for _, sport := range sports {
		for eventIdx := range *sport {
			(*sport)[eventIdx].Team1 = common.ReformatSpaces((*sport)[eventIdx].Team1)
			(*sport)[eventIdx].Team2 = common.ReformatSpaces((*sport)[eventIdx].Team2)
		}
	}

	var knownBM types.Sports
	if err := common.LoadJson(config.FonbetSamplesDir+"knownBookmaker", &knownBM); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(knownBM, *bookmaker) {
		t.Fatal("not equal with known result")
	}
}

func TestBrokenStructure(t *testing.T) {
	data, err := ioutil.ReadFile(config.FonbetSamplesDir + "BrokenStructure.html")
	if err != nil {
		t.Fatal(err)
	}
	html := string(data)
	if _, err := Parse(html); err == nil {
		t.Fatal("no error on broken structure")
	}
}

func TestRealSamples(t *testing.T) {
	for idx := 0; idx < 3; idx++ {
		data, err := ioutil.ReadFile(fmt.Sprintf("%sSample%d.html", config.FonbetSamplesDir, idx))
		if err != nil {
			t.Fatal(err)
		}
		html := string(data)
		if _, err := Parse(html); err != nil {
			t.Fatalf("sample (%d): %v", idx, err)
		}
	}
}
