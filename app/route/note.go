package route

import (
	"encoding/base64"

	"github.com/DSMdongly/glove/app/handler"
	"github.com/DSMdongly/glove/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Note(ech *echo.Echo) {
	ech.POST("/notes", handler.CreateNote)
	ech.GET("/notes", handler.ReadNotes)
	ech.PUT("/notes/:id", handler.UpdateNote)
	ech.DELETE("/notes/:id", handler.DeleteNote)

	not := ech.Group("/notes")

	sec := base64.StdEncoding.EncodeToString([]byte(config.JWT["SECRET"]))

	not.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(sec),
		ContextKey:  "token",
		TokenLookup: "header:Access-Token",
	}))
}
