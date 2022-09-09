package main

import (
	"fmt"
	"time"
)


type User struct {
	FirstName string
	LastName string
	DOB time.Time
}

func (u *User) fullName() string {
	return u.FirstName + u.LastName 
}

func main() {
	user := User{
		FirstName: "Bob",
		LastName: "Hope",
	}

	 fmt.Println(user.fullName())
}
