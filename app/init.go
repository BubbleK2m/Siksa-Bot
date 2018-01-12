package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DSMdongly/siksa-bot/config"
	"github.com/DSMdongly/siksa-bot/support"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	Echo *echo.Echo
)

func Init() {
	Echo = echo.New()
	Echo.Validator = support.NewValidator()

	Echo.Static("/static", "app/static")

	Echo.Use(middleware.Recover())
	Echo.Use(middleware.RequestID())
	Echo.Use(middleware.CORS())
}

func Awake(tik time.Duration) {
	go func(){
		time.Sleep(time.Minute * 1)
		
		for {
			http.Get("https://siksa-bot.herokuapp.com/")
			time.Sleep(tik)
		}
	}()
}

func Start() {
	Echo.Logger.Fatal(Echo.Start(fmt.Sprintf(":%s", config.HTTP["PORT"])))
}
