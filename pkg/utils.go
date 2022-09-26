package pkg

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"

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

func CityToString(c City) string {
	stringValues := []string{
		c.Name,
		c.Region,
		c.District,
		strconv.Itoa(c.Population),
		strconv.Itoa(c.Foundation),
	}
	return strings.Join(stringValues, ",")
}

func CityWithIdToString(c City) string {
	stringValues := []string{
		strconv.Itoa(c.Id),
		c.Name,
		c.Region,
		c.District,
		strconv.Itoa(c.Population),
		strconv.Itoa(c.Foundation),
	}
	return strings.Join(stringValues, ",")
}
