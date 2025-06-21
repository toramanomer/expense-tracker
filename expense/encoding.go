package expense

import (
	"errors"
	"strconv"
	"time"
)

// encode encodes an Expense into a slice of strings.
func encode(expense Expense) []string {
	return []string{
		strconv.Itoa(expense.ID),
		strconv.Itoa(expense.Amount),
		expense.Category,
		expense.Description,
		expense.Date.Format(time.DateOnly),
	}
}

// decode decodes a slice of strings into an [*Expense].
func decode(record []string) (*Expense, error) {
	if len(record) != 5 {
		return nil, errors.New("unexpected record length")
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, errors.New("invalid id: not an integer")
	}

	amount, err := strconv.Atoi(record[1])
	if err != nil {
		return nil, errors.New("invalid amount: not an integer")
	}

	category := record[2]
	description := record[3]
	date, err := time.Parse(time.DateOnly, record[4])
	if err != nil {
		return nil, errors.New("invalid date: not in the format YYYY-MM-DD")
	}

	return &Expense{
		ID:          id,
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        date,
	}, nil
}
