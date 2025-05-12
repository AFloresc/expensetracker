package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {

}

func TestAddExpense(t *testing.T) {
	tracker := Tracker{}
	expense := Expense{
		Id:          0,
		Date:        time.Now(),
		Description: "Lunch",
		Amount:      15,
	}

	tracker.AddExpense(expense)
	fmt.Println(tracker)
	assert.Equal(t, tracker.Elements, 1)
	assert.Equal(t, tracker.Counter, 2)
	assert.Equal(t, tracker.Expenses[0].Id, 1)
	assert.Equal(t, tracker.Expenses[0].Description, expense.Description)

	t.Run("TestListExpenses", func(t *testing.T) {
		_, err := tracker.ListExpenses()
		assert := assert.New(t)

		assert.Nil(err)
	})

	t.Run("TestSummaryExpenses", func(t *testing.T) {
		amount, err := tracker.SummaryExpenses()
		assert := assert.New(t)

		assert.Nil(err)
		if amount != 15.00 {
			t.Errorf("Expected %f, got %f", 15.00, amount)
		}

	})

	t.Run("TestSummaryExpensesByMonth", func(t *testing.T) {
		amount, err := tracker.SummaryExpensesByMonth(5)
		assert := assert.New(t)

		assert.Nil(err)
		if amount != 15.00 {
			t.Errorf("Expected %f, got %f", 15.00, amount)
		}

		amount, err = tracker.SummaryExpensesByMonth(2)
		assert.Nil(err)
		if amount != 0.00 {
			t.Errorf("Expected %f, got %f", 0.00, amount)
		}
	})

	t.Run("TestDeleteExpenseByID", func(t *testing.T) {
		err := tracker.DeleteExpenseByID(1)
		assert := assert.New(t)

		assert.Nil(err)
		if tracker.Elements != 0 {
			t.Errorf("Expected %d, got %d", 0, tracker.Elements)
		}

	})
}

func TestHandleTrackerFile(t *testing.T) {
	tracker := Tracker{}
	expense := Expense{
		Id:          0,
		Date:        time.Now(),
		Description: "Lunch",
		Amount:      15,
	}

	expense2 := Expense{
		Id:          0,
		Date:        time.Now(),
		Description: "Dinner",
		Amount:      15,
	}

	tracker.AddExpense(expense)
	tracker.AddExpense(expense2)
	fmt.Println(tracker)

	err := tracker.HandleTrackerFile("testfile.JSON", "")
	assert := assert.New(t)
	assert.Nil(err)

	err = tracker.SaveTrackerToFile("testfile.JSON")
	assert.Nil(err)
}
