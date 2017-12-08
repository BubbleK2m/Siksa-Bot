package db

import (
	"fmt"

	"github.com/DSMdongly/glove/config"

	mgo "gopkg.in/mgo.v2"
)

var Session *mgo.Session

func Init() (err error) {
	Session, err = mgo.Dial(fmt.Sprintf("%s/%s", config.Mgo["PATH"], config.Mgo["DB"]))

	if err != nil {
		return
	}

	Session.SetSafe(&mgo.Safe{})

	return
}
