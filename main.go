package main

import (
	"log"
	"surebetSearch/chrome"
	"surebetSearch/fonbet"
)

func main() {
	address := "0.0.0.0"
	cdpInfo, err := chrome.Run(address, false)
	if err != nil {
		log.Panic(err)
	}
	defer chrome.Close(cdpInfo)

	if _, err := fonbet.Load(cdpInfo); err != nil {
		log.Panic(err)
	}

	//common.Benchmark(5, []common.FuncInfo{{fonbetNative, "fonbetNative"}})
}
