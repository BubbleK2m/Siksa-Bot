package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "index", nil)
	}
}
