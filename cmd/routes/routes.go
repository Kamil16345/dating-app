package routes

import "net/http"

func routes() {
	http.HandleFunc("/users", users.CreateUser())
}
