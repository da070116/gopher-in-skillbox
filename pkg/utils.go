package pkg

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"io"
	_ "modernc.org/sqlite"
)

const DBFileName string = "database.sqlite3"

// CloseReader - close reader after get request body
func CloseReader(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		logrus.Fatalln(err)
	}
}

// CloseDB - close database
func CloseDB(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Printf("connection to DB  %s closed\n", DBFileName)
}

// CloseQuery - close database
func CloseQuery(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		logrus.Fatalln(err)
	}
}
