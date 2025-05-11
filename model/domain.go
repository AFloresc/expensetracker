package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Expense struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created,omitempty"`
	UpdatedAt   time.Time `json:"updated,omitempty"`
}

type Tracker struct {
	Expenses []Expense `json:"expenses"`
	Counter  int       `json:"counter"`
	Elements int       `json:"elements"`
}

type ExpenseOperations interface {
	AddExpense(e Expense)
	ListExpenses() error
	SummaryExpenses() error
	SummaryExpensesByMonth(month int) (int64, error)
	DeleteExpenseByID(id int) (int64, error)
}

type TrackerOperations interface {
	SaveTrackerToFile(filenaname string) error
	HandleTrackerFile(fileName string, trackerJSON string) error
}

func (tracker *Tracker) AddExpense(e Expense) {
	if tracker.Counter == 0 {
		tracker.Counter++
	}
	e.Id = tracker.Counter
	e.CreatedAt = time.Now()
	tracker.Expenses = append(tracker.Expenses, e)
	tracker.Counter++
	tracker.Elements++
}

func (tracker *Tracker) ListExpenses() (result []Expense, err error) {
	hasElements, err := checkElements(*tracker)
	if hasElements {
		fmt.Println("ID   Date        Description            Amount")
		for _, expense := range tracker.Expenses {
			fmt.Printf("%d    %d-%d-%d    %s                  %.2f€\n", expense.Id, expense.Date.Day(), expense.Date.Month(), expense.Date.Year(), expense.Description, expense.Amount)
		}
	}
	result = tracker.Expenses
	return result, err
}

func (tracker *Tracker) SummaryExpenses() (total float64, err error) {
	hasElements, err := checkElements(*tracker)
	var totalAmount float64 = 0
	if hasElements {

		for _, expense := range tracker.Expenses {
			totalAmount = totalAmount + expense.Amount
		}
		fmt.Printf("Total Expenses: %.2f€\n", totalAmount)
	}
	return totalAmount, err
}

func (tracker *Tracker) SummaryExpensesByMonth(month int) (total float64, err error) {
	var totalAmount float64 = 0
	if month < 1 || month > 12 {
		return totalAmount, errors.New("not a valid month format")
	}
	hasElements, err := checkElements(*tracker)
	if hasElements {

		for _, expense := range tracker.Expenses {
			if expense.Date.Month() == time.Month(month) {
				totalAmount = totalAmount + expense.Amount
			}
		}
		fmt.Printf("Total Expenses: %.2f€\n", totalAmount)
	}
	return totalAmount, err
}

func (tracker *Tracker) DeleteExpenseByID(id int) error {
	hasElements, err := checkElements(*tracker)
	if hasElements {
		for index, expense := range tracker.Expenses {
			if expense.Id == id {
				tracker.Expenses = remove(tracker.Expenses, index)
				tracker.Elements--
				break
			}
		}
	}

	return err
}

func (tracker *Tracker) SaveTrackerToFile(fileName string) error {
	// Tracker to JSON
	data, err := json.MarshalIndent(tracker, "", " ")
	if err != nil {
		return fmt.Errorf("error parsing Event Tracker to JSON: %v", err)
	}

	// Create or write file
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("error writting to file: %v", err)
	}

	fmt.Println("Content saved to file.")
	return nil
}

func (tracker *Tracker) HandleTrackerFile(fileName string, trackerJSON string) error {

	// Check file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// If not created, create it with JSON format
		err := os.WriteFile(fileName, []byte(trackerJSON), 0644)
		if err != nil {
			return fmt.Errorf("error al crear el archivo: %v", err)
		}
		fmt.Println("JSON File created.")
	} else {
		// If created, load tracker content
		data, err := os.ReadFile(fileName)
		if err != nil {
			return fmt.Errorf("error reading file: %v", err)
		}
		err = json.Unmarshal(data, &tracker)
		if err != nil {
			return fmt.Errorf("error unmarshalling file content: %v", err)
		}
		fmt.Println("Content file loaded.")
	}

	return nil
}

func remove(expense []Expense, index int) []Expense {
	return append(expense[:index], expense[index+1:]...)
}

func checkElements(tracker Tracker) (hasElements bool, err error) {
	if tracker.Elements == 0 {
		return false, errors.New("therer are no expenses to list")
	}
	return true, nil
}
