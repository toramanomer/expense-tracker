package expense

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

// Helper function to test adding expenses
func (s *StorageFS) add(expense Expense, w io.Writer) error {
	writer := csv.NewWriter(w)
	if err := writer.Write(encode(expense)); err != nil {
		return err
	}
	writer.Flush()

	return writer.Error()
}

type MyReader struct {
	*strings.Reader
	original string
}

func NewMyReader(s string) *MyReader {
	return &MyReader{
		Reader:   strings.NewReader(s),
		original: s,
	}
}
func (m *MyReader) Truncate(size int64) error {
	if size > m.Size() {
		return errors.New("size is large")
	}
	m.Reset(m.original[0:size])
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

var (
	ErrExpenseNotFound = errors.New("expense not found")
)

type ReadWriteSeekTruncater interface {
	io.ReadWriteSeeker
	Truncate(size int64) error
}

func (s *StorageFS) delete(id int, rw ReadWriteSeekTruncater) error {
	reader := csv.NewReader(rw)
	var (
		records         [][]string
		expenseToDelete *Expense
	)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		expense, err := decode(record)
		if err != nil {
			return err
		}
		if expense.ID == id {
			expenseToDelete = expense
		} else {
			records = append(records, record)
		}
	}

	if expenseToDelete == nil {
		return ErrExpenseNotFound
	}

	if _, err := rw.Seek(0, io.SeekStart); err != nil {
		return err
	}

	if err := rw.Truncate(0); err != nil {
		return err
	}

	// Return early without seeking and writing if there are no records left
	if len(records) == 0 {
		return nil
	}

	return csv.NewWriter(rw).WriteAll(records)
}

// Delete deletes an expense from the file.
// If the file does not exist, [ErrExpenseNotFound] is returned.
func (s *StorageFS) Delete(id int) error {
	file, err := os.OpenFile(s.expensesfile, os.O_RDWR, 0655)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrExpenseNotFound
		}
		return err
	}
	defer file.Close()

	return s.delete(id, file)
}

func (s *StorageFS) list(r io.Reader) ([]Expense, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	expenses := make([]Expense, len(records))
	for i, record := range records {
		expense, err := decode(record)
		if err != nil {
			return nil, err
		}
		expenses[i] = *expense
	}

	return expenses, nil
}

func (s *StorageFS) List() ([]Expense, error) {
	file, err := os.Open(s.expensesfile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()

	return s.list(file)
}

var _ ExpenseStorage = (*StorageFS)(nil)
