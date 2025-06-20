package expense

// StorageFS represents a file system-based storage for expenses.
type StorageFS struct{}

func (s *StorageFS) GenerateID() (int, error) {
	return 0, nil
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
