package main

import (
	"github.com/korovkinand/surebetSearch/dataBase"
	"log"
)

func main() {
	if err := dataBase.Collect(); err != nil {
		log.Panic(err)
	}
}
