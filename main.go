package main

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/database"
)

func main() {

	users := database.NewUsers()
	Bob := users.NewUser("Bob")
	bob := users.NewUser("bob")
	// charlie := users.NewUser("charlie")
	errFol := bob.Follow(Bob)
	errFol = bob.Follow(Bob)
	if errFol != nil {
		fmt.Println(errFol)
	}
	Bob.UploadPhoto()
	_, err := Bob.UploadPhoto()
	if err != nil {
		fmt.Println(err)
	}
	Bob.Like(Bob)

	errli := Bob.Like(Bob)
	if errli != nil {
		fmt.Println(errli)
	}
	td := users.Trending()
	fmt.Println(td)
	fmt.Println(Bob)
	fmt.Println(bob)
	fmt.Println(users)

}
