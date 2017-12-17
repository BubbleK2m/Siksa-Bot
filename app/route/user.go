package route

import (
	"github.com/DSMdongly/glove/app/route/handler"

	"github.com/labstack/echo"
)

func User(ech *echo.Echo) {
	ech.POST("/user/login", handler.Login())
	ech.POST("/user/register", handler.Register())
}
