package expense

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

// Expense represents a single expense entry
type Expense struct {
	ID          int       // Unique identifier for the expense
	Amount      int       // Amount of the expense in dollars
	Category    string    // Category of the expense
	Date        time.Time // Date of the expense
	Description string    // Description of the expense
}

// ValidateID validates the ID of an expense
func ValidateID(id int) error {
	if id <= 0 {
		return errors.New("invalid expense id: must be a positive integer")
	}
	return nil
}

// ValidateAmount validates the amount of an expense
func ValidateAmount(amount int) error {
	if amount <= 0 {
		return errors.New("invalid expense amount: must be a positive integer")
	}
	return nil
}

// ParseCategory parses a category string and returns an error if it's invalid
func ParseCategory(category string) (string, error) {
	category = strings.TrimSpace(category)

	if runeCount := utf8.RuneCountInString(category); runeCount == 0 {
		return "", errors.New("invalid expense category: must not be empty")
	} else if runeCount > 100 {
		return "", errors.New("invalid expense category: must not exceed 100 characters")
	}

	return category, nil
}

// ParseDescription parses a description string and returns an error if it's invalid
func ParseDescription(description string) (string, error) {
	description = strings.TrimSpace(description)

	if runeCount := utf8.RuneCountInString(description); runeCount == 0 {
		return "", errors.New("invalid expense description: must not be empty")
	} else if runeCount > 255 {
		return "", errors.New("invalid expense description: must not exceed 255 characters")
	}

	return description, nil
}
