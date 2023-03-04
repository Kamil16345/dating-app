package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/dating_app")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}
