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
		//assert.Equal(t, len(expenses), int(1))

	})
}
