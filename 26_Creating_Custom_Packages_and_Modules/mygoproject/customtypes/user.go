package customtypes

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func NewUser(firstName, lastName string, age int) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("%s %s (Age: %d)", u.FirstName, u.LastName, u.Age)
}
