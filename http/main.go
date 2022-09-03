package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

// toString - display User data as string
func (u *User) toString() string {
	if len(u.Friends) == 0 {
		return fmt.Sprintf("%s is %d years old. No friends yet\n", u.Name, u.Age)
	}
	return fmt.Sprintf("%s is %d years old. Friends:%v\n", u.Name, u.Age, u.Friends)
}

type service struct {
	storage map[int]User
}

// parseContentToAddFriend - add User by given raw data
func (s *service) parseContentToAddFriend(id int, raw string) (User, error) {
	key, rawFriendId, _ := strings.Cut(raw, "=")
	if key == "friend" {
		friendId, err := strconv.Atoi(rawFriendId)
		if err != nil {
			log.Fatalln(err)
			return User{}, err
		}
		if friendId > len(s.storage) || id > len(s.storage) {
			err := errors.New("no such user")
			return User{}, err
		}

		u := s.storage[id]

		for _, friend := range u.Friends {
			if friendId == friend {
				err := errors.New("are friends already")
				return User{}, err
			}
		}
		u.Friends = append(u.Friends, friendId)
		s.storage[id] = u
		return s.storage[friendId], nil

	} else {
		err := errors.New("bad request")
		log.Fatalln(err)
		return User{}, err
	}
}

// parseContentToAddFriend - add User by given raw data
func (s *service) parseContentToEditAge(id int, raw string) error {
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

		u := s.storage[id]
		u.Age = age
		s.storage[id] = u
		return nil

	} else {
		err := errors.New("bad request")
		log.Fatalln(err)
		return err
	}
}

// Add - add new User record handle
func (s *service) Add(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {

		if request.Header.Get("Content-Type") != "application/json" {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte(errors.New("not a json").Error()))
		}

		content, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		defer closeReader(request.Body)

		var usr User
		if err := json.Unmarshal(content, &usr); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		s.storage[len(s.storage)+1] = usr

		writer.WriteHeader(http.StatusCreated)
		_, _ = writer.Write([]byte(fmt.Sprintf("User was created: %s", usr.toString())))
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		errorString := fmt.Sprintf("method %s not allowed", request.Method)
		_, _ = writer.Write([]byte(errors.New(errorString).Error()))
	}
}

// SetAge - set age for User record
func (s *service) SetAge(writer http.ResponseWriter, request *http.Request) {
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

			if id > len(s.storage) {
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
			status := fmt.Sprintf("%s's age set to %d\n", s.storage[id].Name, s.storage[id].Age)
			_, _ = writer.Write([]byte(status))

		} else {
			writer.WriteHeader(http.StatusBadRequest)
			errorString := "no required query parameter found"
			_, _ = writer.Write([]byte(errors.New(errorString).Error()))
		}
	}
}

// GetList - get all User records handle
func (s *service) GetList(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		if len(s.storage) == 0 {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("No data yet"))
			return
		}
		for id, u := range s.storage {
			_, _ = writer.Write([]byte(fmt.Sprintf("[%d] %v\n", id, u.toString())))
		}
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		errorString := fmt.Sprintf("method %s not allowed", request.Method)
		_, _ = writer.Write([]byte(errors.New(errorString).Error()))
	}
}

// AddFriend - add friend to User
func (s *service) AddFriend(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "PATCH" {
		query := request.URL.Query()
		if query.Has("id") {
			idToAddFriend, err := strconv.Atoi(query.Get("id"))
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

			friend, err := s.parseContentToAddFriend(idToAddFriend, string(content))
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}

			writer.WriteHeader(http.StatusOK)
			status := fmt.Sprintf("%s is %s's friend now", friend.Name, s.storage[idToAddFriend].Name)
			_, _ = writer.Write([]byte(status))

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

// DeleteFromFriendList - remove id from User friends list
func (s *service) DeleteFromFriendList(idToDelete int) {
	for idx, u := range s.storage {
		for i, friendId := range u.Friends {
			if friendId == idToDelete {
				u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
				s.storage[idx] = u
			}
		}
	}
}

// Delete - remove User record handle
func (s *service) Delete(writer http.ResponseWriter, request *http.Request) {
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
			delete(s.storage, idToDelete)
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

// closeReader - close reader after get request body
func closeReader(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// handleRoot - root page handle
func handleRoot(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		errorString := fmt.Sprintf("method %s not allowed", request.Method)
		_, _ = writer.Write([]byte(errors.New(errorString).Error()))
		return
	}
	_, err := writer.Write([]byte(fmt.Sprintf("Server is listening on port %s", request.URL.Port())))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(http.StatusOK)
}

//main - entrypoint for application
func main() {
	fmt.Println("Server instance started")
	st := service{make(map[int]User)}
	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/", handleRoot)
	multiplexer.HandleFunc("/add/", st.Add)
	multiplexer.HandleFunc("/list/", st.GetList)
	multiplexer.HandleFunc("/del/", st.Delete)
	multiplexer.HandleFunc("/friend/", st.AddFriend)
	multiplexer.HandleFunc("/age/", st.SetAge)

	err := http.ListenAndServe("localhost:8080", multiplexer)
	if err != nil {
		log.Fatalln(err)
	}
}
