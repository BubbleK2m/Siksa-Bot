package main

import (
	"log"

	"github.com/DSMdongly/glove/app"
	"github.com/DSMdongly/glove/app/route"
	"github.com/DSMdongly/glove/config"
	"github.com/DSMdongly/glove/db"
)

func main() {
	config.Init()

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	app.Init()
	route.Init()

	app.Start()
}
