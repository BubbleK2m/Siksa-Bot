package handler

import (
	"net/http"

	"github.com/DSMdongly/glove/app/model"

	"github.com/labstack/echo"
)

func Register() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		usr := model.User{}
		db := model.DB

		if err := ctx.Bind(&usr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "binding error")
		}

		if err := ctx.Validate(&usr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "validate error")
		}

		if err := db.Create(&usr).Error; err != nil {
			return echo.NewHTTPError(http.StatusConflict, "account already exist")
		}

		return ctx.NoContent(http.StatusCreated)
	}
}

func Login() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		usr := model.User{}
		db := model.DB

		err := ctx.Bind(&usr)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "binding error")
		}

		err = db.Where("id = ? AND pw = ?", usr.ID, usr.PW).First(&usr).Error

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "failed find user")
		}

		tok, _ := usr.Tokenize()

		return ctx.JSON(http.StatusOK, echo.Map{
			"token": tok,
		})
	}
}
