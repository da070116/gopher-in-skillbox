package pkg

import (
	"errors"
	"net/http"
)

func CustomErrorDisplay(w http.ResponseWriter, status int, errorText string) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(errors.New(errorText).Error()))
}
