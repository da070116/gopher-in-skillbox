package pkg

import "fmt"

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
