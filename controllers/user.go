package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Anirudh-rao/GoMongoRestApi/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

//will Return uc:
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

//Get User is Struct Method
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") //getting ID using Params

	//Checking If Id is HEX
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}

	//Creating A collection of Users in Mongodb and Finding ID
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	//Data Marshling in JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error In Marshelling")
	}
	//Sending response in JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

//function for Creating User
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Creating User Model
	u := models.User{}
	//Decoding the Models
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	//Inserting Models Data
	if err := uc.session.DB("mongo-golang").C("users").Insert(u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error In Marshelling")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

//Function to Delete User
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") //getting ID using Params
	//Checking If Id is HEX
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	//Adding ID to OID
	oid := bson.ObjectIdHex(id)

	//Creating A collection of Users in Mongodb and Finding ID
	if err := uc.session.DB("mongo-golang").C("users").Remove(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "DeletedUser", oid, "\n")

}
