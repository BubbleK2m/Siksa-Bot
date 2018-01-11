package model

import (
	"encoding/base64"

	"github.com/DSMdongly/siksa-bot/config"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	ID string `form:"id" validate:"required" gorm:"primary_key"`
	PW string `form:"pw" validate:"required" gorm:"not_null"`
}

func (usr User) Tokenize() (string, error) {
	tok := jwt.New(jwt.SigningMethodHS256)

	clms := tok.Claims.(jwt.MapClaims)
	clms["id"] = usr.ID

	sec := base64.StdEncoding.EncodeToString([]byte(config.JWT["secret"]))
	str, err := tok.SignedString([]byte(sec))

	if err != nil {
		return "", err
	}

	return str, nil
}
