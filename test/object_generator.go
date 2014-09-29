// +build ignore

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/relops/sqlc/sqlc"
	"os"
)

// TODO paramterize this path
var dbFile = "test/test.db"

func main() {

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		os.Exit(1)
	}

	err = sqlc.Generate(db)
	if err != nil {
		os.Exit(1)
	}
}