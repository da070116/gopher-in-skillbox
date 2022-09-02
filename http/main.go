package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type user struct {
	name    string
	age     int
	friends []int
}

type service struct {
	storage map[int]user
}

// parseContentToAdd - add user by given raw data
func (s *service) parseContentToAdd(raw string) {
	args := strings.Split(raw, "&")
	if len(args) == 2 {
		for i := 0; i < len(args)-1; i++ {
			_, rawName, _ := strings.Cut(args[i], "=")
			name := strings.ReplaceAll(rawName, "%20", " ")
			_, rawAge, _ := strings.Cut(args[i+1], "=")
			age, _ := strconv.Atoi(rawAge)
			userInstance := user{}
			userInstance.name = name
			userInstance.age = age
			uid := len(s.storage) + 1
			s.storage[uid] = userInstance
		}
	}
}

// parseContentToAddFriend - add user by given raw data
func (s *service) parseContentToAddFriend(id int, raw string) {
	key, rawFriendId, _ := strings.Cut(raw, "=")
	if key == "friend" {
		println(key)
		friendId, err := strconv.Atoi(rawFriendId)
		if err != nil {
			log.Fatalln(err)
		}
		u := s.storage[id]
		u.friends = append(u.friends, friendId)
		s.storage[id] = u
	} else {
		err := errors.New("bad request")
		log.Fatalln(err)
	}

}

// Add - add new user record handle
func (s *service) Add(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {

		content, err := ioutil.ReadAll(request.Body)
		if err != nil {
			println("can't read content")
		}

		s.parseContentToAdd(string(content))
		writer.WriteHeader(http.StatusCreated)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// GetList - get all user records handle
func (s *service) GetList(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		if len(s.storage) == 0 {
			_, _ = writer.Write([]byte("No data yet"))
		}
		for id, u := range s.storage {
			_, _ = writer.Write([]byte(fmt.Sprintf("[%d] User named %s is %d years old\n %#v\n", id, u.name, u.age, u.friends)))

		}
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// AddFriend - add friend to user
func (s *service) AddFriend(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "PATCH" {
		query := request.URL.Query()
		if query.Has("id") {
			idToAddFriend, err := strconv.Atoi(query.Get("id"))
			if err != nil {
				log.Fatalln(err)
			}

			content, err := ioutil.ReadAll(request.Body)
			if err != nil {
				println("can't read content")
			}
			println(idToAddFriend, string(content))
			s.parseContentToAddFriend(idToAddFriend, string(content))

		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}

	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Delete - remove user record handle
func (s *service) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "DELETE" {
		query := request.URL.Query()

		if query.Has("id") {
			idToDelete, err := strconv.Atoi(query.Get("id"))
			if err != nil {
				log.Fatalln(err)
			}

			for idx, u := range s.storage {
				for i, friendId := range u.friends {
					if friendId == idToDelete {
						u.friends = append(u.friends[:i], u.friends[i+1:]...)
						s.storage[idx] = u
					}
				}
			}

			delete(s.storage, idToDelete)
			writer.WriteHeader(http.StatusNoContent)
			_, _ = writer.Write([]byte("record deleted"))
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// handleRoot - endpoint to get root page
func handleRoot(writer http.ResponseWriter, request *http.Request) {

	_, err := writer.Write([]byte("Server is listening: " + strconv.FormatInt(time.Now().UnixNano(), 10)))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)
}

//main - entrypoint for application
func main() {
	fmt.Println("Server instance started")
	st := service{make(map[int]user)}
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", handleRoot)
	multiplexer.HandleFunc("/add/", st.Add)
	multiplexer.HandleFunc("/list/", st.GetList)
	multiplexer.HandleFunc("/del/", st.Delete)
	multiplexer.HandleFunc("/friend/", st.AddFriend)
	err := http.ListenAndServe("localhost:8080", multiplexer)
	if err != nil {
		log.Fatalln(err)
	}

}
