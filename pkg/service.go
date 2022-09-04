package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Service struct {
	Storage map[int]User
}

// parseContentToAddFriend - add User by given raw data
func (s *Service) parseContentToAddFriend(id int, raw string) (User, error) {
	key, rawFriendId, _ := strings.Cut(raw, "=")
	if key == "friend" {
		friendId, err := strconv.Atoi(rawFriendId)
		if err != nil {
			log.Fatalln(err)
			return User{}, err
		}
		if friendId > len(s.Storage) || id > len(s.Storage) {
			err := errors.New("no such user")
			return User{}, err
		}

		u := s.Storage[id]

		for _, friend := range u.Friends {
			if friendId == friend {
				err := errors.New("are friends already")
				return User{}, err
			}
		}
		u.Friends = append(u.Friends, friendId)
		s.Storage[id] = u
		return s.Storage[friendId], nil

	} else {
		err := errors.New("bad request")
		log.Fatalln(err)
		return User{}, err
	}
}

// parseContentToAddFriend - add User by given raw data
func (s *Service) parseContentToEditAge(id int, raw string) error {
	key, rawAge, _ := strings.Cut(raw, "=")
	if key == "age" {
		age, err := strconv.Atoi(rawAge)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		if age < 0 {
			err := errors.New("wrong age was given")
			return err
		}

		u := s.Storage[id]
		u.Age = age
		s.Storage[id] = u
		return nil

	} else {
		err := errors.New("bad request")
		log.Fatalln(err)
		return err
	}
}

// Add - add new User record handle
func (s *Service) Add(writer http.ResponseWriter, content []byte) {

	var usr User
	if err := json.Unmarshal(content, &usr); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	s.Storage[len(s.Storage)+1] = usr

	writer.WriteHeader(http.StatusCreated)
	_, _ = writer.Write([]byte(fmt.Sprintf("User was created: %s", usr.toString())))
}

// SetAge - set age for User record
func (s *Service) SetAge(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "PATCH" {
		query := request.URL.Query()
		if query.Has("id") {
			id, err := strconv.Atoi(query.Get("id"))
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			content, err := ioutil.ReadAll(request.Body)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			defer closeReader(request.Body)

			if id > len(s.Storage) {
				err := errors.New("no such user")
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}

			err = s.parseContentToEditAge(id, string(content))
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}

			writer.WriteHeader(http.StatusOK)
			status := fmt.Sprintf("%s's age set to %d\n", s.Storage[id].Name, s.Storage[id].Age)
			_, _ = writer.Write([]byte(status))

		} else {
			writer.WriteHeader(http.StatusBadRequest)
			errorString := "no required query parameter found"
			_, _ = writer.Write([]byte(errors.New(errorString).Error()))
		}
	}
}

// GetList - get all User records handle
func (s *Service) GetList(writer http.ResponseWriter) {

	if len(s.Storage) == 0 {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("No data yet"))
		return
	}
	for id, u := range s.Storage {
		_, _ = writer.Write([]byte(fmt.Sprintf("[%d] %v\n", id, u.toString())))
	}
	writer.WriteHeader(http.StatusOK)

}

// AddFriend - add friend to User
func (s *Service) AddFriend(writer http.ResponseWriter, idToAddFriend int, content string) {
	friend, err := s.parseContentToAddFriend(idToAddFriend, content)
	if err != nil {
		CustomErrorDisplay(writer, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	status := fmt.Sprintf("%s is %s's friend now", friend.Name, s.Storage[idToAddFriend].Name)
	_, _ = writer.Write([]byte(status))

}

// DeleteFromFriendList - remove id from User friends list
func (s *Service) DeleteFromFriendList(idToDelete int) {
	for idx, u := range s.Storage {
		for i, friendId := range u.Friends {
			if friendId == idToDelete {
				u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
				s.Storage[idx] = u
			}
		}
	}
}

// Delete - remove User record handle
func (s *Service) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "DELETE" {
		query := request.URL.Query()

		if query.Has("id") {
			idToDelete, err := strconv.Atoi(query.Get("id"))
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}

			s.DeleteFromFriendList(idToDelete)
			delete(s.Storage, idToDelete)
			writer.WriteHeader(http.StatusNoContent)
			_, _ = writer.Write([]byte("record deleted"))
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			errorString := "no required query parameter found"
			_, _ = writer.Write([]byte(errors.New(errorString).Error()))
		}
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		errorString := fmt.Sprintf("method %s not allowed", request.Method)
		_, _ = writer.Write([]byte(errors.New(errorString).Error()))
	}
}
