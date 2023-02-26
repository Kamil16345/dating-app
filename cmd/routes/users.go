package routes

import (
	"database/sql"
	"fmt"
	"net/http"
)
func CreateUser(w http.ResponseWriter, *r http.Request) {
	w.WriteHeader(400)

	username := r.FormValue("username")
	email := r.FormValue("email")

	result, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
	if err != nil {
		http.Error(w, "Error inserting user.", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Fprintf(w, "Inserted %d rows into users table.", rowsAffected)
}