package main

import (
	"fmt"
	".example/struct/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser *User

	// strcut literal
	// appUser = User{
	// 	firstName: userFirstName,
	// 	lastName: userLastName,
	// 	birtDate: userBirthdate,
	// 	createdAt: time.Now(),

	// 	// alternatif
	// 	// bisa langsung define variable-nya, tapi harus berurutan
	// }

	// Cara make struct di folder lain
	// appUser, err := &user.User{
	// 	FirstName: userFirstName
	// }

	// pake constructor
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Print(err)
		return 
	}

	admin := user.NewAdmin("a@gmail.com", "123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data!

	// fmt.Println(firstName, lastName, birthdate)
	// bisa pake pointer disini incase struct-nya expand terlalu besar
	appUser.OutputUserDetails()
	// fmt.Print(appUser)

	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	// pake scanln biar enter diartikan sebagai value
	fmt.Scanln(&value)
	return value
}