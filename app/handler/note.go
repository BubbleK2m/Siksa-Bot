package handler

import (
	"net/http"

	"github.com/DSMdongly/glove/app/model"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func CreateNote(ctx echo.Context) error {
	not := model.Note{}

	if err := ctx.Bind(&not); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "binding error")
	}

	if err := ctx.Validate(&not); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validate error")
	}

	if err := not.Create(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot create note")
	}

	return ctx.NoContent(http.StatusCreated)
}

func ReadNotes(ctx echo.Context) error {
	nots, err := model.Note{}.FindAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed read notes")
	}

	return ctx.JSON(http.StatusOK, nots)
}

func UpdateNote(ctx echo.Context) error {
	id := bson.ObjectIdHex(ctx.Param("id"))

	not, err := model.Note{}.FindByID(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "cannot find note")
	}

	if err := ctx.Bind(&not); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "binding error")
	}

	if err := ctx.Validate(&not); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := not.Update(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed update note")
	}

	return ctx.NoContent(http.StatusOK)
}

func DeleteNote(ctx echo.Context) error {
	id := bson.ObjectIdHex(ctx.Param("id"))

	not, err := model.Note{}.FindByID(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "cannot find note")
	}

	if err := not.Delete(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed delete note")
	}

	return ctx.NoContent(http.StatusOK)
}
