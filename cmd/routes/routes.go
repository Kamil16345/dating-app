package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Router = mux.NewRouter()

func Routes() {

	http.Handle("/", Router)
	Router.HandleFunc("/users", GetUsers).Methods("GET")
	Router.HandleFunc("/users/:id", GetUser).Methods("GET")
	Router.HandleFunc("/users", CreateUser).Methods("POST")
	Router.HandleFunc("/users/:id", UpdateUser).Methods("PUT")
	Router.HandleFunc("/users/:id", DeleteUser).Methods("DELETE")

}
