package api

import (
	mysql "dating-app/internal/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mysql.ConnectDB()
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
