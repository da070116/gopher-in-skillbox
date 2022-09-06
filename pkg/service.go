package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Database struct {
	Conn *sql.DB
}

// Add - add new User record handle
func (db *Database) Add(writer http.ResponseWriter, content []byte) {

	var usr User
	if err := json.Unmarshal(content, &usr); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}

	insertSQL := `INSERT INTO users(name, age) VALUES (?, ?)`
	statement, err := db.Conn.Prepare(insertSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(usr.Name, usr.Age)
	if err != nil {
		log.Fatalln(err.Error())
	}

	writer.WriteHeader(http.StatusCreated)
	_, _ = writer.Write([]byte(fmt.Sprintf("User was created: %s", usr.toString())))
}

// SetAge - set age for User record
func (db *Database) SetAge(writer http.ResponseWriter, userID int, ageValue int) {

	querySetAge := `UPDATE users SET age = ?  WHERE id = ?`

	stmt, err := db.Conn.Prepare(querySetAge)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = stmt.Exec(ageValue, userID)
	if err != nil {
		log.Fatalln(err.Error())
	}
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("Age was changed"))
}

// GetList - get all User records handle
func (db *Database) GetList(writer http.ResponseWriter) {

	selectSQL := `SELECT * FROM users;`
	rows, err := db.Conn.Query(selectSQL)
	if err != nil {
		if err == sql.ErrNoRows {
			_, _ = writer.Write([]byte("No records yet"))
			writer.WriteHeader(http.StatusOK)
			return
		} else {
			log.Fatalln(err.Error())
		}
	}
	defer CloseQuery(rows)

	getFriendsSQL := `SELECT friend_id FROM friends WHERE owner_id=?`
	stmt, err := db.Conn.Prepare(getFriendsSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		_ = rows.Scan(&id, &name, &age)
		var outputStr string
		friends, err := stmt.Query(id)
		if err != nil {
			log.Fatalln(err.Error())
		}
		var friendList []string
		for friends.Next() {
			var friendId string
			_ = friends.Scan(&friendId)
			friendList = append(friendList, friendId)

		}

		friendsDisplay := strings.Join(friendList, ",")
		outputStr = fmt.Sprintf("[%d]: %s is %d years old. Friends: [%s]\n", id, name, age, friendsDisplay)
		_, _ = writer.Write([]byte(outputStr))
		CloseQuery(friends)
	}
	writer.WriteHeader(http.StatusOK)
}

// AddFriend - add friend to User
func (db *Database) AddFriend(writer http.ResponseWriter, idToAddFriend int, friendID int) {

	if idToAddFriend == friendID {
		http.Error(writer, "can not add the same user", http.StatusBadRequest)
		return
	}

	var availableIDs []int

	checkSQL := `SELECT id FROM users;`
	ids, err := db.Conn.Query(checkSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for ids.Next() {
		var availableID int
		_ = ids.Scan(&availableID)
		availableIDs = append(availableIDs, availableID)
	}
	CloseQuery(ids)

	var first, second bool
	for id := range availableIDs {
		if idToAddFriend == id {
			first = true
		}
		if friendID == id {
			second = true
		}
	}

	if first == false || second == false {
		http.Error(writer, "can not add user", http.StatusBadRequest)
		return
	}

	insertSQL := `INSERT INTO friends(owner_id, friend_id) VALUES (?, ?)`
	statement, err := db.Conn.Prepare(insertSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(idToAddFriend, friendID)
	if err != nil {
		log.Fatalln(err.Error())
	}

	writer.WriteHeader(http.StatusCreated)
	_, _ = writer.Write([]byte("friend was added"))

}

// DeleteFromFriendList - remove id from User friends list
func (db *Database) DeleteFromFriendList(idToDelete int) {
	queryDelete := `DELETE FROM friends WHERE owner_id = ? OR friend_id = ?`

	stmt, err := db.Conn.Prepare(queryDelete)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = stmt.Exec(idToDelete, idToDelete)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Delete - remove User record handle
func (db *Database) Delete(writer http.ResponseWriter, idToDelete int) {

	queryDelete := `DELETE FROM users WHERE id = ?`
	stmt, err := db.Conn.Prepare(queryDelete)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = stmt.Exec(idToDelete)
	if err != nil {
		log.Fatalln(err.Error())
	}

	writer.WriteHeader(http.StatusNoContent)
	_, _ = writer.Write([]byte("Record was deleted"))
}
