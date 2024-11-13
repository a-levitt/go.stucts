package main

import (
	"alevitt.com/go.structs/user"
	"fmt"
)

func main() {
	userFirstName := user.GetUserData("Please enter your first name: ")
	userLastName := user.GetUserData("Please enter your last name: ")
	userBirthDate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin, _ := user.NewAdmin("artemis.levitt@gmail.com", "testpass")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data!

	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()

	appUser.ClearUserName()
	appUser.OutputUserDetails()
}
