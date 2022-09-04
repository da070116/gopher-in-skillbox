package pkg

import (
	"io"
	"log"
)

// closeReader - close reader after get request body
func closeReader(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
