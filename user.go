package gopher_in_skillbox

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type UpdateUserData struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

type UserFriendData struct {
	FriendId *int `json:"friendId"`
}
