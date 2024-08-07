package main

import (
	"example.com/structs/user"
	_ "example.com/structs/user"
	"fmt"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str
	name = "Jimmy"
	name.log()

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "passwording")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data!

	// outputUserDetails(&appUser) // pointer is not good use here but a practice
	appUser.OutputUserDetails() // method way
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		return ""
	}
	return value
}
