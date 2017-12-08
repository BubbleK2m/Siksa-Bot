package handler

import (
	"net/http"

	"github.com/DSMdongly/glove/app/model"

	"github.com/labstack/echo"
)

func Register(ctx echo.Context) error {
	usr := model.User{}

	if err := ctx.Bind(&usr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "binding error")
	}

	if err := ctx.Validate(&usr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validate error")
	}

	if err := usr.Create(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed create note")
	}

	return ctx.NoContent(http.StatusCreated)
}

func Login(ctx echo.Context) error {
	usr := model.User{}

	err := ctx.Bind(&usr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "binding error")
	}

	usr, err = model.User{}.FindByIDAndPW(usr.ID, usr.PW)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "failed find user")
	}

	tok, _ := usr.Tokenize()

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": tok,
	})
}
