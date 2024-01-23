package main

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/database"
)

func main() {
	users := database.NewUsers()
	bob := users.NewUser("bob")
	Bob := users.NewUser("Bob")
	bob.Follow(Bob)
	errFol := bob.Follow(Bob)
	if errFol != nil {
		fmt.Println(errFol)
	}
	Bob.UploadPhoto()
	_, err := Bob.UploadPhoto()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Bob)
	fmt.Println(bob)
	fmt.Println(users)
}
