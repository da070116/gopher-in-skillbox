package pkg

import (
	"io"
	"log"
)

// CloseReader - close reader after get request body
func CloseReader(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
