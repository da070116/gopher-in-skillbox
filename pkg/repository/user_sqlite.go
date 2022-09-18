package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	gopherinskillbox "skillbox-test"
	"skillbox-test/pkg"
	"strings"
)

type UserSqlite struct {
	db *sqlx.DB
}

func NewUserSqlite(db *sqlx.DB) *UserSqlite {
	return &UserSqlite{db: db}
}

// CreateUser database query function
func (r *UserSqlite) CreateUser(user gopherinskillbox.User) (gopherinskillbox.User, error) {
	createUserQuery := `INSERT INTO users(name, age) VALUES (?, ?)`
	stmt, err := r.db.Prepare(createUserQuery)
	if err != nil {
		return gopherinskillbox.User{}, err
	}

	_, err = stmt.Exec(user.Name, user.Age)
	if err != nil {
		return gopherinskillbox.User{}, err
	}
	return user, err
}

// GetAllUsers database query function
func (r *UserSqlite) GetAllUsers() ([]gopherinskillbox.User, error) {
	getUsersQuery := `SELECT * FROM users;`
	usersResult, err := r.db.Query(getUsersQuery)
	if err != nil {
		return nil, err
	}
	usersList := make([]gopherinskillbox.User, 0)
	for usersResult.Next() {
		var (
			userID, userAge int
			userName        string
		)

		err = usersResult.Scan(&userID, &userName, &userAge)
		if err != nil {
			return nil, err
		}

		// select friends for each id
		getUserFriendsQuery := `SELECT friend_id FROM friends WHERE owner_id = ?;`

		friendsResult, err := r.queryStatement(getUserFriendsQuery, userID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		var friendsIdList []int
		for friendsResult.Next() {
			var friendID int
			err := friendsResult.Scan(&friendID)
			if err != nil {
				return nil, err
			}
			friendsIdList = append(friendsIdList, friendID)
		}
		usersList = append(usersList, gopherinskillbox.User{
			Name:    userName,
			Age:     userAge,
			Friends: friendsIdList,
		})
	}
	return usersList, nil
}

// DeleteUser database query function
func (r *UserSqlite) DeleteUser(deleteId int) error {
	deleteFriendsQuery := `DELETE FROM friends WHERE friend_id = $1 OR owner_id = $1`
	err := r.executeStatement(deleteFriendsQuery, deleteId)
	if err != nil {
		return err
	}

	deleteSqlQuery := `DELETE FROM users WHERE id = ?;`

	return r.executeStatement(deleteSqlQuery, deleteId)
}

// UpdateUser database query function
func (r *UserSqlite) UpdateUser(updateId int, userData gopherinskillbox.UpdateUserData) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if userData.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *userData.Name)
		argId++
	}
	if userData.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *userData.Age)
		argId++
	}
	args = append(args, updateId)
	alterValues := strings.Join(setValues, ", ")
	fmtString := "UPDATE users SET" + " " + alterValues + " WHERE id=%d"
	updateQuery := fmt.Sprintf(fmtString, updateId)
	_, err := r.db.Exec(updateQuery, args...)
	return err
}

// AddFriend database query function
func (r *UserSqlite) AddFriend(id int, userFriendData gopherinskillbox.UserFriendData) error {

	if id == *userFriendData.FriendId {
		return errors.New("can't add user itself as friend")
	}

	checkIdsQuery := `SELECT id FROM users`
	rows, err := r.db.Query(checkIdsQuery)
	if err != nil {
		return err
	}
	existedIds := make([]int, 0)
	for rows.Next() {
		var existedId int
		err := rows.Scan(&existedId)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
				break
			} else {
				return err
			}
		}
		existedIds = append(existedIds, existedId)
	}
	pkg.CloseQuery(rows)

	var ownerIdMatch, friendIdMatch bool
	for _, val := range existedIds {
		if val == id {
			ownerIdMatch = true
		}
		if val == *userFriendData.FriendId {
			friendIdMatch = true
		}
	}
	if !friendIdMatch || !ownerIdMatch {
		return errors.New("related id not found")
	}

	existedFriendsQuery := `SELECT * FROM friends WHERE owner_id = ? AND friend_id = ?`
	stmt, err := r.db.Prepare(existedFriendsQuery)
	existedFriends, err := stmt.Query(id, *userFriendData.FriendId)
	if err != nil {
		return err
	}
	if existedFriends.Next() {
		return errors.New("are friends already")
	}
	pkg.CloseQuery(existedFriends)

	addFriendQuery := `INSERT INTO friends (owner_id, friend_id) VALUES (?, ?)`
	stmt, err = r.db.Prepare(addFriendQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, userFriendData.FriendId)
	return err
}

// executeStatement - execute prepared change query
func (r *UserSqlite) executeStatement(query string, id int) error {
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// queryStatement - execute prepared select query
func (r *UserSqlite) queryStatement(query string, id int) (*sql.Rows, error) {
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	friendsResult, err := stmt.Query(id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer pkg.CloseQuery(friendsResult)
	return friendsResult, nil
}
