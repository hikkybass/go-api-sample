package infra

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectDB() *sql.DB {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	return db
}
