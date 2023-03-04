package routes

import (
	"database/sql"
	"dating-app/cmd/models"
	mysql "dating-app/internal/db"
	"encoding/json"
	"log"
	"net/http"
)

var db *sql.DB

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	db := mysql.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err = db.Exec("INSERT INTO users(name, email, password) VALUES(?, ?, ?)", name, email, password)
	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
