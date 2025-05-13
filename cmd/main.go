package main

import (
	"expensetracker/model"
	"fmt"
	"os"
	"strconv"
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
		summaryEvents(tracker, argWithProg)
	case "delete":
		fmt.Println("list received")
		deleteEvent(tracker, argWithProg)
	case "":
		fmt.Println("no needed")

	}
}

func deleteEvent(tracker model.Tracker, argWithProg []string) {
	if len(argWithProg) != 4 {
		fmt.Println(fmt.Errorf("bad arguments").Error())
	}
	if argWithProg[3] != "--id" {
		fmt.Println(fmt.Errorf("bad arguments").Error())
	}
	expenseid, err := strconv.Atoi(argWithProg[3])
	if err != nil {
		fmt.Println(err)
	}
	err = tracker.DeleteExpenseByID(expenseid)
	if err != nil {
		fmt.Println(err)
	}
	err = tracker.SaveTrackerToFile(JSONFileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Event deleted id: ", expenseid)
}

func summaryEvents(tracker model.Tracker, argWithProg []string) {
	var totalAmount float64 = 0.00
	if len(argWithProg) == 2 {
		total, err := tracker.SummaryExpenses()
		if err != nil {
			fmt.Println(err)
		}
		totalAmount = total
	}
	if len(argWithProg) == 3 {
		fmt.Println(fmt.Errorf("bad arguments").Error())
	}
	if len(argWithProg) == 4 {
		if argWithProg[3] != "--month" {
			fmt.Println(fmt.Errorf("bad arguments").Error())
		}
		month, err := strconv.Atoi(argWithProg[3])
		if err != nil {
			fmt.Println(err)
		}
		total, err := tracker.SummaryExpensesByMonth(month)
		if err != nil {
			fmt.Println(err)
		}
		totalAmount = total
	}
	fmt.Printf("Total Expenses: %.2f€\n", totalAmount)
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

func validateArgs(argPos int, args []string) bool {
	if len(args) < argPos+1 {
		return false
	}
	return true
}
