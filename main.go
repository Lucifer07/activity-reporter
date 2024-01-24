package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/activityreporter"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/utility"
)

func promptInput(scanner *bufio.Scanner, text string) []string {
	fmt.Print(text)
	scanner.Scan()
	textInput := scanner.Text()
	inputSplit := strings.Split(textInput, " ")
	var inpUser []string
	for _, input := range inputSplit {
		if input != "" {
			inpUser = append(inpUser, input)
		}
	}
	return inpUser
}
func main() {
	start := true
	reader := bufio.NewScanner(os.Stdin)
	users := activityreporter.NewUsers()
	menu := "\nActivity Reporter\n" +
		"1. Setup\n" +
		"2. Action\n" +
		"3. Display\n" +
		"4. Trending\n" +
		"5. Exit"
	for start {
		fmt.Println(menu)
		dashboardMenu := promptInput(reader, "Enter menu: ")
		switch dashboardMenu[0] {
		case "1":
			setupMenu := promptInput(reader, "Setup social graph: ")
			if len(setupMenu) != 3 || setupMenu[1] != "follows" {
				fmt.Println(utility.ErrKeyword.Error())
				break
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
		case "2":
			actionMenu := promptInput(reader, "Enter user Actions: ")
			if len(actionMenu) == 3 || len(actionMenu) == 4 {
				if len(actionMenu) == 3 {
					if (actionMenu[1] != "uploaded") || (actionMenu[2] != "photo") {
						fmt.Println(utility.ErrKeyword.Error())
						break
					}
					user, err := users.CheckUser(actionMenu[0])
					if err != nil {
						fmt.Println(err.Error())
						break
					}
					_, err = user.UploadPhoto()
					if err != nil {
						fmt.Println(err.Error())
						break
					}
					break
				}
				if actionMenu[1] != "likes" && actionMenu[3] != "photo" {
					fmt.Println(utility.ErrKeyword.Error())
					break
				}
				firstPerson, err := users.CheckUser(actionMenu[0])
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				secondPerson, err := users.CheckUser(actionMenu[2])
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				err = firstPerson.Like(secondPerson)
				if err != nil {
					fmt.Println(err.Error())
				}
				break
			}
			fmt.Println(utility.ErrKeyword.Error())
		case "3":
			displayMenu := promptInput(reader, "Display activity for: ")
			user, err := users.CheckUser(displayMenu[0])
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("\n" + user.Name() + " activities:")
			for _, activity := range user.Activity() {
				fmt.Println(activity)
			}
		case "4":
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
		case "5":
			fmt.Println("Good bye!")
			start = false
			break
		default:
			fmt.Println("invalid menu")
		}
	}
}
