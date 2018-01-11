package main

import (
	"log"

	"github.com/DSMdongly/siksa-bot/app"
	"github.com/DSMdongly/siksa-bot/app/model"
	"github.com/DSMdongly/siksa-bot/app/route"
	"github.com/DSMdongly/siksa-bot/config"
)

func main() {
	config.Init()
	app.Init()

	if err := model.Init(); err != nil {
		log.Fatal(err)
	}

    route.Init()
	
	bot.Init()
	bot.Start()

	app.Start()
}
