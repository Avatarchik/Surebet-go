package main

import (
	"github.com/korovkinand/surebetSearch/config"
	"github.com/korovkinand/surebetSearch/db"
	"log"
)

func main() {
	if err := db.Collect(config.DbPath); err != nil {
		log.Panic(err)
	}
}
