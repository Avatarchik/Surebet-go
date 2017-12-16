package main

import (
	"github.com/korovkinand/surebetSearch/scanners/db"
	"log"
)

func main() {
	if err := db.Collect(); err != nil {
		log.Panic(err)
	}
}
