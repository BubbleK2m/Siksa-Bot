package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Alarm struct {
	gorm.Model

	Note string    `json:"note" validate:"required" gorm:"not_null" form:"note"`
	Time time.Time `json:"time" validate:"required" gorm:"not_null;unique_index"`
}
