package user

import (
	"errors"
	"fmt"
	"time"
)

// Semua penamaan harus capital kalau mau di export

type User struct {
	firstName string
	lastName string
	birtDate string
	createdAt time.Time
}

// STRUCT EMBED
type Admin struct {
	User
	email string
	password string
}

// jadiin method
func (data User) OutputUserDetails() {
	
	// normalnya, harus di dereference kalau pake tipe data lain
	// khusus struct, boleh pake shortcut langsung tanpa dereferencing 
	fmt.Println(data.firstName, data.lastName, data.birtDate)
}

func (data *User) ClearUserName() {
	data.firstName = ""
	data.lastName = ""
}

// constructor function
// bisa pake pointer in case struct complex
// Untuk penamaan-nya, normalnya cuma pake New
func New(firstName, lastName, birthDate string) (*User, error) {

	if len(firstName) < 1 || len(lastName) < 1 || birthDate == ""{
		return nil, errors.New("First name, last name, or birthdate required")
	}

	return &User {
		firstName: firstName,
		lastName: lastName,
		birtDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return  Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birtDate: "---",
			createdAt: time.Now(),
		},
	}
}