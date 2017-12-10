package main

import (
	"log"

	"glove/app"
	"glove/app/model"
	"glove/app/route"
	"glove/config"
)

func main() {
	config.Init()
	app.Init()

	if err := model.Init(); err != nil {
		log.Fatal(err)
	}

	route.Init()
	app.Start()
}
