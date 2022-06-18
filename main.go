package main

import (
	"net/http"

	"github.com/Anirudh-rao/GoMongoRestApi/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	//Creating Our Basic Router
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	//Routes
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/:id", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	//creating a Local Host
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
