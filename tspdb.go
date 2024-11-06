package tspdb

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkErrorFatal(err error, exit bool) {
	if err != nil {
		log.Fatal(err)
		if exit {
			os.Exit(1)
		}
	}
}

func InitDatabaseConnection(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	checkError(err)

	 _, err = db.Exec("CREATE TABLE IF NOT EXISTS example (id INTEGER PRIMARY KEY, name TEXT)") checkError(err)
	checkError(err)
	return db, nil
}
