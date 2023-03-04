package main

import (
	"dating-app/cmd/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes.Routes()
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", routes.Router))
}
