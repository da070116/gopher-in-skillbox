package pkg

import (
	"database/sql"
	"io"
	"log"
	_ "modernc.org/sqlite"
)

const DBFileName string = "database.sqlite3"

// CloseReader - close reader after get request body
func CloseReader(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// CloseDB - close database
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func ConfigureDatabase() {

	db, err := sql.Open("sqlite", DBFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer CloseDB(db)

	createTable(db)
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE user(
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" VARCHAR(200),
    "age" INTEGER NOT NULL,
    "friends" TEXT
);`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	}
}
