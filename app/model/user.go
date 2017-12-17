package model

import (
	"encoding/base64"

	"github.com/DSMdongly/glove/config"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	ID string `json:"id" form:"id" validate:"required" gorm:"primary_key"`
	PW string `json:"pw" form:"pw" validate:"required"`
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
