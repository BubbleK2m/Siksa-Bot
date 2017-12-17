package main

import (
	"log"

	"github.com/DSMdongly/glove/app"
	"github.com/DSMdongly/glove/app/doc"
	"github.com/DSMdongly/glove/app/model"
	"github.com/DSMdongly/glove/app/route"
	"github.com/DSMdongly/glove/config"
)

func main() {
	config.Init()
	app.Init()

	if err := model.Init(); err != nil {
		log.Fatal(err)
	}

	doc.Init()
	route.Init()

	app.Start()
}
