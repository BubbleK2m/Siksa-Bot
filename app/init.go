package app

import (
	"fmt"

	"github.com/DSMdongly/glove/config"
	"github.com/DSMdongly/glove/support"
	"github.com/savaki/swag"
	"github.com/savaki/swag/swagger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	Echo    *echo.Echo
	Swagger *swagger.API
)

func Init() {
	Swagger = swag.New(
		swag.Title("Glove API Document"),
		swag.Description("API document for Glove"),
		swag.Tag("user", "API about user"),
	)

	Echo = echo.New()
	Echo.Validator = support.NewValidator()

	Echo.Static("/static", "app/static")

	Echo.Use(middleware.Recover())
	Echo.Use(middleware.RequestID())
	Echo.Use(middleware.CORS())
}

func Start() {
	Echo.Logger.Fatal(Echo.Start(fmt.Sprintf(":%s", config.HTTP["PORT"])))
}
