package app

import (
	"fmt"

	"github.com/DSMdongly/glove/config"
	"github.com/DSMdongly/glove/support"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	validator "gopkg.in/go-playground/validator.v9"
)

var Echo *echo.Echo

func Init() {
	Echo = echo.New()
	Echo.Validator = support.Validator{Validation: validator.New()}

	Echo.Use(middleware.Recover())
	Echo.Use(middleware.RequestID())
	Echo.Use(middleware.CORS())
}

func Start() {
	Echo.Logger.Fatal(Echo.Start(fmt.Sprintf(":%s", config.HTTP["PORT"])))
}
