package route

import (
	"github.com/DSMdongly/siksa-bot/app/route/handler"

	"github.com/labstack/echo"
)

func User(ech *echo.Echo) {
	ech.POST("/user/auth", handler.Auth())
	ech.POST("/user/register", handler.Register())
}
