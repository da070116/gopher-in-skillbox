package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Service struct {
	Storage map[int]User
}

// parseContentToAddFriend - add User by given raw data
func (s *Service) parseContentToAddFriend(userID int, friendID int) (User, error) {

	if friendID > len(s.Storage) || userID > len(s.Storage) {
		err := errors.New("no such user")
		return User{}, err
	}

	u := s.Storage[userID]

	if friendID == userID {
		err := errors.New("same ids")
		return User{}, err
	}

	for _, friend := range u.Friends {
		if friendID == friend {
			err := errors.New("are friends already")
			return User{}, err
		}
	}
	u.Friends = append(u.Friends, friendID)
	s.Storage[userID] = u
	return s.Storage[friendID], nil

}

// parseContentToAddFriend - add User by given raw data
func (s *Service) parseContentToEditAge(id int, ageVal int) error {

	if ageVal < 0 {
		err := errors.New("wrong age was given")
		return err
	}

	u := s.Storage[id]
	u.Age = ageVal
	s.Storage[id] = u
	return nil
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
func (s *Service) SetAge(writer http.ResponseWriter, userID int, ageValue int) {

	err := s.parseContentToEditAge(userID, ageValue)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
		return
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
func (s *Service) AddFriend(writer http.ResponseWriter, idToAddFriend int, friendID int) {
	friend, err := s.parseContentToAddFriend(idToAddFriend, friendID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	status := fmt.Sprintf("%s is %s's friend now", friend.Name, s.Storage[idToAddFriend].Name)
	_, _ = writer.Write([]byte(status))

}

// DeleteFromFriendList - remove id from User friends list
func (s *Service) DeleteFromFriendList(idToDelete int) {
	for idx, u := range s.Storage {
		for i, friendID := range u.Friends {
			if friendID == idToDelete {
				u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
				s.Storage[idx] = u
			}
		}
	}
}

// Delete - remove User record handle
func (s *Service) Delete(writer http.ResponseWriter, idToDelete int) {

	delete(s.Storage, idToDelete) // if no such id - just ignore
	writer.WriteHeader(http.StatusNoContent)
	_, _ = writer.Write([]byte("record deleted"))
}
