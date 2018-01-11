package route

import (
	"encoding/base64"

	"github.com/DSMdongly/siksa-bot/app/route/handler"
	"github.com/DSMdongly/siksa-bot/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Alarm(ech *echo.Echo) {
	sec := base64.StdEncoding.EncodeToString([]byte(config.JWT["secret"]))

	alm := ech.Group("/alarm")
	alm.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		AuthScheme: "JWT",
		SigningKey: []byte(sec),
	}))

	ech.POST("/alarm", handler.CreateAlarm())
	ech.GET("/alarm", handler.ReadAlarms())
	ech.DELETE("/alarm/:id", handler.DeleteAlarm())
}
