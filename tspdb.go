package github.com/flutesnz/tspdb

/*
	- create/init
	- update
	- delete
	- view

*/

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

func initDatabaseConnection(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	checkError(err)
	return db, nil
}

func main() {
	db, dbError := initDatabaseConnection("./tsp.db")
	checkErrorFatal(dbError, true)
	defer db.Close()
}
