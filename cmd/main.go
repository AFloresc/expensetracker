package main

import (
	"expensetracker/model"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const JSONFileName = "trackData.JON"
const instructions = "Possible commands: \n - ?: Shows this help \n - summary: Shows the sum of all expenses\n - summary --month <month_number>: Shows the sum of al the expneses in selected month\n - delete --id <expense_id>: Deletes the selected expense with id\n - list: shows a list of all the expenses\n"
const version = "V. 0.1"

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
		addEvent(tracker, argWithProg)
	case "list":
		fmt.Println("list received")
		listEvents(tracker)
	case "summary":
		// Can have --month argument
		fmt.Println("list received")
		summaryEvents(tracker, argWithProg)
	case "delete":
		// Must have --id argument
		fmt.Println("list received")
		deleteEvent(tracker, argWithProg)
	case "?":
		fmt.Print(instructions)
	case "version":
		fmt.Println(version)
	case "":
		fmt.Print("BAd arguments \n" + instructions)

	}
}

func addEvent(tracker model.Tracker, argWithProg []string) {
	if len(argWithProg) != 6 {
		fmt.Println(fmt.Errorf("bad arguments").Error())
	}
	if argWithProg[3] != "--description" {
		fmt.Println(fmt.Errorf("bad arguments").Error())
	}
	if argWithProg[5] != "--amount" {
		fmt.Println(fmt.Errorf("bad arguments").Error())
	}

	amount, err := strconv.ParseFloat(argWithProg[6], 32)
	if err != nil {
		fmt.Println(err)
	}
	expense := model.Expense{
		Date:        time.Now(),
		Description: argWithProg[4],
		Amount:      amount,
	}

	tracker.AddExpense(expense)
	err = tracker.SaveTrackerToFile(JSONFileName)
	if err != nil {
		fmt.Println(err)
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
