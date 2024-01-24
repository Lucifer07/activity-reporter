package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/activityreporter"
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
			firstCase(users, setupMenu)
		case "2":
			actionMenu := promptInput(reader, "Enter user Actions: ")
			secondCase(users, actionMenu)
		case "3":
			displayMenu := promptInput(reader, "Display activity for: ")
			thirdCase(users, displayMenu)
		case "4":
			fourthCase(users)
		case "5":
			fmt.Println("Good bye!")
			start = false
			break
		default:
			fmt.Println("invalid menu")
		}
	}
}
