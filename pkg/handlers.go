package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// HandleUserGetList - show all Users
func HandleUserGetList(writer http.ResponseWriter, request *http.Request, s *Service) {
	if request.Method == "GET" {
		s.GetList(writer)
	} else {
		errorText := fmt.Sprintf("method %s is not allowed", request.Method)
		CustomErrorDisplay(writer, http.StatusMethodNotAllowed, errorText)
	}
}

// HandleUserAddNew - add new User
func HandleUserAddNew(writer http.ResponseWriter, request *http.Request, s *Service) {
	if request.Method == "POST" {

		if request.Header.Get("Content-Type") != "application/json" {
			CustomErrorDisplay(writer, http.StatusBadRequest, "not a JSON")
			return
		}

		content, err := ioutil.ReadAll(request.Body)
		if err != nil {
			CustomErrorDisplay(writer, http.StatusInternalServerError, err.Error())
			return
		}
		defer closeReader(request.Body)

		s.Add(writer, content)

	} else {
		errorText := fmt.Sprintf("method %s is not allowed", request.Method)
		CustomErrorDisplay(writer, http.StatusMethodNotAllowed, errorText)
	}
}

// HandleUserAddNewFriend - add new friend in list for User by id
func HandleUserAddNewFriend(writer http.ResponseWriter, request *http.Request, s *Service) {
	if request.Method == "PATCH" {
		query := request.URL.Query()
		if query.Has("id") {
			idToAddFriend, err := strconv.Atoi(query.Get("id"))
			if err != nil {
				CustomErrorDisplay(writer, http.StatusBadRequest, err.Error())
				return
			}
			content, err := ioutil.ReadAll(request.Body)
			if err != nil {
				CustomErrorDisplay(writer, http.StatusInternalServerError, err.Error())
				return
			}
			defer closeReader(request.Body)

			s.AddFriend(writer, idToAddFriend, string(content))

		} else {
			CustomErrorDisplay(writer, http.StatusBadRequest, "no required query parameter found")
		}

	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		errorString := fmt.Sprintf("method %s not allowed", request.Method)
		_, _ = writer.Write([]byte(errors.New(errorString).Error()))
	}
}

// HandleUserSetAge - set age for User by id
func HandleUserSetAge() {

}

// HandleUserDelete - delete User by id
func HandleUserDelete() {

}
