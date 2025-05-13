package main

import (
	"expensetracker/model"
	"fmt"
	"os"
	"strings"
)

const JSONFileName = "trackData.JON"

func main() {
	argWithProg := os.Args

	if len(argWithProg) < 2 {
		fmt.Println("Plesa add the right arguments ")
		return
	}

	mainArgument := strings.ToLower(argWithProg[1])
	tracker, err := loadTrackerFromFile(JSONFileName)
	if err != nil {
		fmt.Println(err)
	}

	switch mainArgument {
	case "add":
		fmt.Println("add received")
	case "list":
		fmt.Println("list received")
		listEvents(tracker)
	case "summary":
		//Can have --month, --id argument or none
		fmt.Println("list received")
	case "delete":
		fmt.Println("list received")
	case "":
		fmt.Println("no needed")

	}
}

func listEvents(tracker model.Tracker) {
	expenses, err := tracker.ListExpenses()
	if err != nil {
		fmt.Println(err)
	}
	if len(expenses) > 0 {
		fmt.Println("ID   Date        Description            Amount")
		for _, expense := range tracker.Expenses {
			fmt.Printf("%d    %d-%d-%d    %s                  %.2f€\n", expense.Id, expense.Date.Day(), expense.Date.Month(), expense.Date.Year(), expense.Description, expense.Amount)
		}
	}
}

func loadTrackerFromFile(fileName string) (model.Tracker, error) {
	var tracker model.Tracker

	err := tracker.HandleTrackerFile(fileName, "")
	if err != nil {
		return tracker, err
	}
	return tracker, err
}
