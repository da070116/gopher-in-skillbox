package pkg

import (
	"database/sql"
	"io"
	"log"
	"os"
)

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
	f, err := os.Create("database.db")
	if err != nil {
		log.Fatalln(err)
	}
	_ = f.Close()

	db, _ := sql.Open("sqlite3", "database.db")
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
