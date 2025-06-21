package expense

import (
	"os"
	"strconv"
)

// StorageFS represents a file system-based storage for expenses.
type StorageFS struct {
	idsfile string
}

// GenerateID generates a new unique ID for an expense starting from 1
// Creates a new file if it doesn't exist, otherwise truncates and writes the new ID.
func (s *StorageFS) GenerateID() (int, error) {
	data, err := os.ReadFile(s.idsfile)
	// Only return error if it's not a file not found error
	if err != nil && !os.IsNotExist(err) {
		return 0, err
	}

	var lastID int
	if len(data) > 0 {
		id, err := strconv.Atoi(string(data))
		if err != nil {
			return 0, err
		}
		lastID = id
	}

	newID := lastID + 1
	if err := os.WriteFile(s.idsfile, []byte(strconv.Itoa(newID)), os.ModePerm); err != nil {
		return 0, err
	}

	return newID, nil
}

func (s *StorageFS) Add(expense Expense) error {
	return nil
}

func (s *StorageFS) Delete(id int) error {
	return nil
}

func (s *StorageFS) List() ([]Expense, error) {
	return nil, nil
}

var _ ExpenseStorage = (*StorageFS)(nil)
