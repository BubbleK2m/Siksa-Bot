package handler

import (
	"net/http"
	"time"

	"github.com/DSMdongly/siksa-bot/app/model"
	"github.com/DSMdongly/siksa-bot/support"

	"github.com/labstack/echo"
)

func CreateAlarm() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		db := model.DB
		alm := model.Alarm{}

		if err := ctx.Bind(&alm); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "binding error")
		}

		if tme := ctx.FormValue("time"); tme == "" {
			alm.Time = time.Now()
		} else {
			alm.Time, _ = support.ParseTime(tme)
		}

		alm.Time = time.Date(2000, 6, 8, alm.Time.Hour(), alm.Time.Minute(), alm.Time.Second(), 3, alm.Time.Location())

		if err := ctx.Validate(&alm); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "validation error")
		}

		if err := db.Create(&alm).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to create alarm")
		}

		return ctx.JSON(http.StatusCreated, echo.Map{
			"created": alm.ID,
		})
	}
}

func ReadAlarms() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		db := model.DB
		alms := []model.Alarm{}

		if err := db.Find(&alms).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "cannot found alarms")
		}

		return ctx.JSON(http.StatusOK, alms)
	}
}

func DeleteAlarm() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		db := model.DB
		alm := model.Alarm{}

		if err := db.Where("id = ?", ctx.Param("id")).First(&alm).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "cannot found alarm by id")
		}

		if err := db.Delete(&alm).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete alram")
		}

		return ctx.JSON(http.StatusOK, echo.Map{
			"deleted": alm.ID,
		})
	}
}
