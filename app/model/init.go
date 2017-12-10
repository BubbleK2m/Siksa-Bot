package model

import (
	"glove/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Init() (err error) {
	DB, err = gorm.Open("sqlite3", config.Sqlite["PATH"])

	if err != nil {
		return
	}

	if err = DB.AutoMigrate(&User{}).Error; err != nil {
		return
	}

	return nil
}
