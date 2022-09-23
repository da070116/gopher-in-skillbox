package pkg

import (
	"encoding/json"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// CloseReader - close reader after get request body
func CloseReader(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		logrus.Fatalln(err)
	}
}

func CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		logrus.Fatalln(err)
	}
}

func JSONFormatString(uid int, c City) (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		logrus.Fatalln(err)
	}
	return string(data), nil
}
