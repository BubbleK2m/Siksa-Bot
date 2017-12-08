package model

import (
	"encoding/base64"
	"time"

	"github.com/DSMdongly/glove/config"
	"github.com/DSMdongly/glove/db"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID string `json:"id" bson:"id" form:"id" validate:"required"`
	PW string `json:"pw" bson:"pw" form:"pw" validate:"required"`
}

func (usr *User) Create() error {
	ses := db.Session.Clone()
	defer ses.Close()

	col := ses.DB(config.Mgo["DB"]).C("users")
	err := col.Insert(*usr)

	return err
}

func (User) FindByIDAndPW(id, pw string) (usr User, err error) {
	ses := db.Session.Clone()
	defer ses.Close()

	col := ses.DB(config.Mgo["DB"]).C("users")
	qry := bson.M{
		"id": id,
		"pw": pw,
	}

	if err = col.Find(qry).One(&usr); err != nil {
		return
	}

	return
}

func (usr User) Tokenize() (str string, err error) {
	tok := jwt.New(jwt.SigningMethodHS256)
	clms := tok.Claims.(jwt.MapClaims)

	clms["id"] = usr.ID
	clms["exp"] = time.Now().Add(time.Hour * 72).Unix()

	sec := base64.StdEncoding.EncodeToString([]byte(config.JWT["SECRET"]))

	str, err = tok.SignedString([]byte(sec))

	return
}
