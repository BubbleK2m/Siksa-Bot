package route

import (
	"github.com/DSMdongly/glove/app/handler"

	"github.com/labstack/echo"
)

func User(ech *echo.Echo) {
	ech.POST("/login", handler.Login)
	ech.POST("/register", handler.Register)
}
