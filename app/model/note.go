package model

import (
	"github.com/DSMdongly/glove/config"
	"github.com/DSMdongly/glove/db"

	"gopkg.in/mgo.v2/bson"
)

type Note struct {
	ID      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string        `json:"title" bson:"title" form:"title" validate:"required"`
	Content string        `json:"content" bson:"content" form:"content" validate:"required"`
}

func (not Note) Create() error {
	ses := db.Session.Clone()
	defer ses.Close()

	col := ses.DB(config.Mgo["DB"]).C("notes")
	err := col.Insert(not)

	return err
}

func (Note) FindAll() (nots []Note, err error) {
	ses := db.Session.Clone()
	defer ses.Close()

	col := ses.DB(config.Mgo["DB"]).C("notes")

	if err = col.Find(nil).All(&nots); err != nil {
		return
	}

	return
}

func (Note) FindByID(id bson.ObjectId) (not Note, err error) {
	ses := db.Session.Clone()
	defer ses.Close()

	col := ses.DB(config.Mgo["DB"]).C("notes")
	qry := bson.M{
		"_id": id,
	}

	if err = col.Find(qry).One(&not); err != nil {
		return
	}

	return
}

func (not Note) Update() error {
	ses := db.Session.Clone()
	defer ses.Close()

	col := ses.DB(config.Mgo["DB"]).C("notes")
	err := col.Update(bson.M{"_id": not.ID}, not)

	return err
}

func (not Note) Delete() error {
	ses := db.Session.Clone()
	defer ses.Close()

	col := ses.DB(config.Mgo["DB"]).C("notes")
	err := col.Remove(bson.M{"_id": not.ID})

	return err
}
