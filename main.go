package main

import (
	"log"
	"surebetSearch/chrome"
	"surebetSearch/fonbet"
	"github.com/knq/chromedp/client"
	"fmt"
)

func printTargets(cdpInfo *chrome.CDPInfo, targetClient *client.Client, msg string)  {
	log.Printf("LIST OF TARGETS %s:", msg)
	targets, err := targetClient.ListTargets(cdpInfo.Ctxt)
	if err != nil{
		log.Panic(err)
	}
	log.Println("CLient: ")
	for _, target := range targets{
		log.Println(target)
	}
	log.Println("CDP: ")
	for _, target := range cdpInfo.C.ListTargets(){
		log.Println(target)
	}
}

func main() {
	address := "0.0.0.0"
	cdpInfo, err := chrome.Run(address, false)
	if err != nil {
		log.Panic(err)
	}
	defer chrome.Close(cdpInfo)

	targetClient := client.New()

	printTargets(cdpInfo, targetClient, "Initially")

	for i := 0 ; i < 3; i++{
		if err := fonbet.Load(cdpInfo); err != nil{
			log.Panic(err)
		}
		printTargets(cdpInfo, targetClient, fmt.Sprintf("After %d new tab", i+1))
	}

	//common.Benchmark(5, []common.FuncInfo{{fonbetNative, "fonbetNative"}})
}
