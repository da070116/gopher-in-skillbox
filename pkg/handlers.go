package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//HandlersManager - base function for navigation.
//It defines all available paths for REST API and checks
//whether request.URL.Path matches one of them
//
//Available paths and methods are:
//	   '/users/' - (GET) - display list of all users
//	   '/users/' - (POST) - Add new User
//	'/users/id/' - (PATCH) - edit User age or add friends to list (depends on parameter in request body)
//	'/users/id/' - (DELETE) - remove User
//
func HandlersManager(writer http.ResponseWriter, request *http.Request, service *Service) {

	if request.URL.Path == "/users/" {

		switch request.Method {
		case http.MethodGet:
			handleUserGetList(writer, service)
			break
		case http.MethodPost:
			handleUserAddNew(writer, request, service)
			break
		default:
			s := fmt.Sprintf("method %s not supported", request.Method)
			http.Error(writer, s, http.StatusNotImplemented)
			return
		}

	} else if strings.HasPrefix(request.URL.Path, "/users/") {

		path := strings.Trim(request.URL.Path, "/") // clear path string from `/` at begin and end
		pathParts := strings.Split(path, "/")

		// if path didn't contain id
		if len(pathParts) < 2 {
			http.Error(writer, "expect <id> in this path", http.StatusBadRequest)
			return
		}

		// obtain User id from path
		id, err := strconv.Atoi(pathParts[1])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		// and check if there is a User with this id
		if id > len(service.Storage) {
			http.Error(writer, "no such user", http.StatusBadRequest)
			return
		}

		switch request.Method {

		case http.MethodDelete:
			handleUserDelete(writer, service, id)
			break

		case http.MethodPatch:
			handleUserPatch(writer, request, service, id)

		default:
			s := fmt.Sprintf("method %s not supported", request.Method)
			http.Error(writer, s, http.StatusNotImplemented)
			return
		}

	} else {
		errorText := fmt.Sprintf("Handler %s is not supported", request.URL.Path)
		http.Error(writer, errorText, http.StatusNotImplemented)
		return
	}
}

// handleUserGetList - show all Users
func handleUserGetList(writer http.ResponseWriter, s *Service) {
	s.GetList(writer)
}

// handleUserDelete - delete a User
func handleUserDelete(writer http.ResponseWriter, s *Service, userID int) {
	s.DeleteFromFriendList(userID)
	s.Delete(writer, userID)
}

// handleUserAddNew - add new User
func handleUserAddNew(writer http.ResponseWriter, request *http.Request, s *Service) {

	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer closeReader(request.Body)
	s.Add(writer, content)
}

// handleUserPatch - alter User data (age or friends list)
func handleUserPatch(writer http.ResponseWriter, request *http.Request, s *Service, userID int) {
	// validate request body
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	key, rawValue, err := parseBody(string(content))
	switch key {
	case "friend":
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			s := fmt.Sprintf("value of %s parameter should be integer", key)
			http.Error(writer, s, http.StatusBadRequest)
			return
		}

		s.AddFriend(writer, userID, value)
		break
	case "age":
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			s := fmt.Sprintf("value of %s parameter should be integer", key)
			http.Error(writer, s, http.StatusBadRequest)
			return
		}
		s.SetAge(writer, userID, value)
		break
	default:
		s := fmt.Sprintf("parameter %s not supported", key)
		http.Error(writer, s, http.StatusBadRequest)
		return
	}
}

// parseBody - split request body with single parameter
func parseBody(raw string) (key string, value string, err error) {
	key, value, found := strings.Cut(raw, "=")
	if found {
		return key, value, nil
	} else {
		return "", "", errors.New("bad request")
	}
}
