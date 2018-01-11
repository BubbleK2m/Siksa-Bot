package handler

import (
	"net/http"

	"github.com/DSMdongly/siksa-bot/app/model"
	"github.com/labstack/echo"
)

func Auth() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		db := model.DB
		usr := model.User{}

		if err := ctx.Bind(&usr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "binding error")
		}

		if err := db.Where("id = ? and pw = ?", usr.ID, usr.PW).First(&usr).Error; err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "cannot find user")
		}

		tok, _ := usr.Tokenize()

		return ctx.JSON(http.StatusOK, echo.Map{
			"token": tok,
		})
	}
}

func Register() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		db := model.DB
		usr := model.User{}

		if err := ctx.Bind(&usr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "binding error")
		}

		if err := ctx.Validate(&usr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "validation error")
		}

		if err := db.Create(&usr).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
		}

		tok, _ := usr.Tokenize()

		return ctx.JSON(http.StatusCreated, echo.Map{
			"token": tok,
		})
	}
}
