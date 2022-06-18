package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	//We are specifying the User Datatypes in Our Project
	//So we are specifing what is Looks in JSON and
	//For Mongo What it  Looks as Bson
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"Name" bson:"Name"`
	Gender string        `json:"Gender" bson:"Gender"`
	Age    int           `json:"Age" bson:"Age"`
}
