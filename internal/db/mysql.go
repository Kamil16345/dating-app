package mysql

import (
	"database/sql"
	"fmt"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	fmt.Println("Success, connected to DB!")
	return db, nil
}
