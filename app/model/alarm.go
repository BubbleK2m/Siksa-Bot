package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Alarm struct {
	gorm.Model

	Note string    `validate:"required" gorm:"not_null" form:"note"`
	Time time.Time `validate:"required" gorm:"not_null;unique_index"`
}
