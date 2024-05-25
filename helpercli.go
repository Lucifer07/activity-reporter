package main

import (
	"fmt"
	"strconv"

	"github.com/Lucifer07/activity-reporter/activityreporter"
	"github.com/Lucifer07/activity-reporter/utility"
)

func firstCase(users *activityreporter.Users, setupMenu []string) {
	if len(setupMenu) != 3 || setupMenu[1] != "follows" {
		fmt.Println(utility.ErrKeyword.Error())
		return
	}
	firstUser, firstErr := users.CheckUser(setupMenu[0])
	if firstErr != nil {
		firstUser = users.NewUser(setupMenu[0])
	}
	secondUser, secondErr := users.CheckUser(setupMenu[2])
	if secondErr != nil {
		secondUser = users.NewUser(setupMenu[2])
	}
	err := firstUser.Follow(secondUser)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func secondCase(users *activityreporter.Users, actionMenu []string) {
	if len(actionMenu) == 3 || len(actionMenu) == 4 {
		if len(actionMenu) == 3 {
			handleUploadedArgument(users, actionMenu)
			return
		}
		handleLikesArgument(users, actionMenu)
		return
	}
	fmt.Println(utility.ErrKeyword.Error())
	return
}
func thirdCase(users *activityreporter.Users, displayMenu []string) {
	user, err := users.CheckUser(displayMenu[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n" + user.Name() + " activities:")
	for _, activity := range user.Activity() {
		fmt.Println(activity)
	}
	return
}
func fourthCase(users *activityreporter.Users) {
	fmt.Println("\nTrending photos:")
	for index, photo := range users.Trending() {
		sentence := ""
		if photo.Like != 1 {
			sentence = strconv.Itoa(index+1) + ". " + photo.Name + " photo got " + strconv.Itoa(photo.Like) + " likes"
		} else {
			sentence = strconv.Itoa(index+1) + ". " + photo.Name + " photo got " + strconv.Itoa(photo.Like) + " like "
		}
		fmt.Println(sentence)
	}
}
func handleUploadedArgument(users *activityreporter.Users, actionMenu []string) {
	if (actionMenu[1] != "uploaded") || (actionMenu[2] != "photo") {
		fmt.Println(utility.ErrKeyword.Error())
		return
	}
	user, err := users.CheckUser(actionMenu[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = user.UploadPhoto()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
func handleLikesArgument(users *activityreporter.Users, actionMenu []string) {
	if actionMenu[1] != "likes" || actionMenu[3] != "photo" {
		fmt.Println(utility.ErrKeyword.Error())
		return
	}
	firstPerson, err := users.CheckUser(actionMenu[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	secondPerson, err := users.CheckUser(actionMenu[2])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = firstPerson.Like(secondPerson)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
