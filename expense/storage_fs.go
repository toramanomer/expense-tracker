package expense

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// StorageFS represents a file system-based storage for expenses.
type StorageFS struct {
	idsfile      string
	expensesfile string
}

// NewStorageFS creates a new StorageFS instance.
// It initializes the storage directory, panics if it fails due to an error other than file already exists.
func NewStorageFS(dirname string) *StorageFS {
	if err := os.MkdirAll(dirname, os.ModePerm); err != nil && !os.IsExist(err) {
		panic(err)
	}

	return &StorageFS{
		idsfile:      filepath.Join(dirname, "ids.txt"),
		expensesfile: filepath.Join(dirname, "expenses.txt"),
	}
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

type F interface {
	io.Writer
}

func (s *StorageFS) add(expense Expense, w io.Writer) error {
	writer := csv.NewWriter(w)
	record := []string{
		strconv.Itoa(expense.ID),
		expense.Description,
		strconv.Itoa(expense.Amount),
		expense.Date.Format(time.DateOnly),
	}
	if err := writer.Write(record); err != nil {
		return err
	}
	writer.Flush()

	return nil
}

func (s *StorageFS) Add(expense Expense) error {
	file, err := os.OpenFile(s.expensesfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0655)
	if err != nil {
		return err
	}
	defer file.Close()

	return s.add(expense, file)
}

func (s *StorageFS) Delete(id int) error {
	return nil
}

func (s *StorageFS) List() ([]Expense, error) {
	return nil, nil
}

var _ ExpenseStorage = (*StorageFS)(nil)
