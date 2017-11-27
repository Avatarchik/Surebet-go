package main

import (
	"log"
	"surebetSearch/fonbet"
	"surebetSearch/common"
	"sync"
	"context"
	"github.com/knq/chromedp"
	"surebetSearch/chrome"
)


func main() {
	address := "0.0.0.0"

	targetNumber := 10

	tabsLoad := func() error{
		cdpInfo, err := chrome.Run(address, false)
		if err != nil {
			log.Panic(err)
		}
		defer chrome.Close(cdpInfo)

		for i := 0; i < targetNumber; i++ {
			if err := fonbet.Load(cdpInfo); err != nil {
				return err
			}
			//chrome.PrintTargets(cdpInfo)
			var html string
			if err := cdpInfo.C.Run(cdpInfo.Ctxt, chrome.GetHtml(&html)); err != nil {
				return err
			}
			log.Printf("Html size: %d", len(html))
		}
		return nil
	}

	poolLoad := func() error{
		ctxt, cancel := context.WithCancel(context.Background())
		defer cancel()

		port := 9222
		// create pool
		pool, err := chromedp.NewPool(chromedp.PortRange(port, port + targetNumber + 1))
		if err != nil {
			return err
		}
		// loop over the URLs
		var wg sync.WaitGroup
		for i := 0; i < targetNumber; i++{
			wg.Add(1)
			go fonbet.LoadPool(ctxt, &wg, pool, i)
		}
		// wait for to finish
		wg.Wait()

		// shutdown pool
		err = pool.Shutdown()
		if err != nil {
			return err
		}
		return nil
	}

	if err := common.Benchmark(1, []common.FuncInfo{
		{poolLoad, "poolLoad"},
		{tabsLoad, "tabsLoad"},
		}); err != nil{
			log.Panic(err)
	}
}
