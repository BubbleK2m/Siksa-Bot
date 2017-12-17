package route

import (
	"github.com/DSMdongly/glove/app"
	"github.com/labstack/echo"
)

func Doc(ech *echo.Echo) {
	app.Echo.GET("/swagger", echo.WrapHandler(app.Swagger.Handler(true)))
}
