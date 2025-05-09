package model

import (
	"errors"
	"fmt"
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
	SummaryExpensesByMonth(month int) error
	DeleteExpenseByID(id int) error
}

func (tracker *Tracker) AddExpense(e Expense) {
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
			fmt.Printf("%d   %d-%d-%d        %s            %f\n", expense.Id, expense.Date.Day(), expense.Date.Month(), expense.Date.Year(), expense.Description, expense.Amount)
		}
	}
	result = tracker.Expenses
	return result, err
}

func (tracker *Tracker) SummaryExpenses() (err error) {
	hasElements, err := checkElements(*tracker)
	if hasElements {
		var totalAmount float64 = 0
		for _, expense := range tracker.Expenses {
			totalAmount = totalAmount + expense.Amount
		}
		fmt.Println("Total Expenses: ", totalAmount)
	}
	return err
}

func (tracker *Tracker) SummaryExpensesByMonth(month int) (err error) {
	if month < 1 || month > 12 {
		return errors.New("not a valid month format")
	}
	hasElements, err := checkElements(*tracker)
	if hasElements {
		var totalAmount float64 = 0
		for _, expense := range tracker.Expenses {
			if expense.Date.Month() == time.Month(month) {
				totalAmount = totalAmount + expense.Amount
			}
		}
		fmt.Println("Total Expenses: ", totalAmount)
	}
	return err
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

func remove(expense []Expense, index int) []Expense {
	return append(expense[:index], expense[index+1:]...)
}

func checkElements(tracker Tracker) (hasElements bool, err error) {
	if tracker.Elements == 0 {
		return false, errors.New("therer are no expenses to list")
	}
	return true, nil
}
