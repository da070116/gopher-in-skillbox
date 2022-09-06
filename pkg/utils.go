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
	println("connection to DB closed")
}

// CloseQuery - close database
func CloseQuery(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// ConfigureDatabase - create db file and launch initialization
func ConfigureDatabase() {
	db, err := sql.Open("sqlite", DBFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer CloseDB(db)
	createTables(db)

}

// createTables - create tables in database
func createTables(db *sql.DB) {
	query := `CREATE TABLE users(
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" VARCHAR(200),
    "age" INTEGER NOT NULL
);`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	}

	query = `CREATE TABLE friends(
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "owner_id" INTEGER NOT NULL, 
        "friend_id" INTEGER NOT NULL,
        CONSTRAINT "fk_owner" 
        	foreign key ("owner_id")
            references users(id), 
        CONSTRAINT "fk_friend" 
        	foreign key ("friend_id")
            references users(id)
		);`

	stmt, err = db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	}
}

// DatabaseConnection - init connection
func DatabaseConnection() (db *sql.DB) {
	db, err := sql.Open("sqlite", DBFileName)
	if err != nil {
		log.Fatalln(err)
	}
	println("connected to " + DBFileName)
	return
}
