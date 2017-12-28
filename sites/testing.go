package sites

import (
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/surebetSearch/chrome"
	"github.com/korovkinand/surebetSearch/common"
	"time"
)

func testLoad(s *common.SiteInfo, initLoad, loadEvents chromedp.Action) error {
	if err := chrome.RunPool(1); err != nil {
		return err
	}
	defer chrome.ClosePool()

	if err := chrome.RunActions(initLoad); err != nil {
		return err
	}

	var actions chromedp.Tasks
	testNum := 1

	if loadEvents != nil {
		actions = append(actions, loadEvents)
		testNum = 3
	}

	var html string
	actions = append(actions,
		chrome.SaveScn(s.Url),
		chromedp.OuterHTML(s.Node, &html),
	)

	for test := 0; test < testNum; test++ {
		if err := chrome.RunActions(actions); err != nil {
			return err
		}

		if len(html) == 0 {
			common.Error("empty html")
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func TestLoadEvents(s *common.SiteInfo, initLoad, loadEvents chromedp.Action) error {
	return testLoad(s, initLoad, loadEvents)
}

func TestInitLoad(s *common.SiteInfo, initLoad chromedp.Action) error {
	return testLoad(s, initLoad, nil)
}
