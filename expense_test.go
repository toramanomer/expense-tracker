package expense

import (
	"errors"
	"sync"
	"testing"
)

type mockStorage struct {
	id        int
	idErr     error
	addErr    error
	deleteErr error
	listErr   error

	mu       sync.Mutex
	expenses []Expense
}

func newMockStorage() *mockStorage {
	return &mockStorage{}
}

func (m *mockStorage) GenerateID() (int, error) {
	return m.id, m.idErr
}

func (m *mockStorage) Add(expense Expense) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.addErr != nil {
		return m.addErr
	}

	m.expenses = append(m.expenses, expense)
	return nil
}

func (m *mockStorage) Delete(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.expenses {
		if e.ID == id {
			m.expenses = append(m.expenses[:i], m.expenses[i+1:]...)
			return nil
		}
	}
	return m.deleteErr
}

func (m *mockStorage) List() ([]Expense, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.listErr != nil {
		return nil, m.listErr
	}

	return m.expenses, nil
}

func TestExpenseService_AddExpense(t *testing.T) {
	t.Run("fails with failed id generation", func(t *testing.T) {
		s := newMockStorage()
		s.idErr = errors.New("gen id err")
		service := ExpenseService{expenseStorage: s}
		expense, err := service.AddExpense("desc", 10)

		if err == nil {
			t.Error("expected error, got none")
		}
		if err != nil && !errors.Is(err, s.idErr) {
			t.Errorf("unexpected error: %v", err)
		}

		if expense != nil {
			t.Errorf("expected expense to be nil, got: %v", expense)
		}
	})

	t.Run("fails with failed add", func(t *testing.T) {
		s := newMockStorage()
		s.addErr = errors.New("add err")
		service := ExpenseService{expenseStorage: s}
		expense, err := service.AddExpense("desc", 10)

		if err == nil {
			t.Error("expected error, got none")
		}
		if err != nil && !errors.Is(err, s.addErr) {
			t.Errorf("unexpected error: %v", err)
		}

		if expense != nil {
			t.Errorf("expected expense to be nil, got: %v", expense)
		}
	})

	t.Run("successful add", func(t *testing.T) {
		s := newMockStorage()
		s.id = 3
		service := ExpenseService{expenseStorage: s}
		expense, err := service.AddExpense("desc", 10)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if expense == nil {
			t.Fatal("expected expense to be not nil")
		}

		if expense.ID != s.id {
			t.Errorf("expected expense ID to be %d, got: %d", s.id, expense.ID)
		}

		if expense.Description != "desc" {
			t.Errorf("expected expense description to be %s, got: %s", "desc", expense.Description)
		}

		if expense.Amount != 10 {
			t.Errorf("expected expense amount to be %d, got: %d", 10, expense.Amount)
		}

		if expense.Date.IsZero() {
			t.Errorf("expected expense date to be not zero, got: %v", expense.Date)
		}
	})
}

func TestExpenseService_DeleteExpense(t *testing.T) {
	t.Run("fails with storage err", func(t *testing.T) {
		s := newMockStorage()
		s.deleteErr = errors.New("delete err")
		service := ExpenseService{expenseStorage: s}
		err := service.DeleteExpense(1)

		if err == nil {
			t.Error("expected error, got none")
		}
		if err != nil && !errors.Is(err, s.deleteErr) {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestExpenseService_ListExpenses(t *testing.T) {
	t.Run("fails with storage err", func(t *testing.T) {
		s := newMockStorage()
		s.listErr = errors.New("list err")
		service := ExpenseService{expenseStorage: s}
		expenses, err := service.ListExpenses()

		if err == nil {
			t.Error("expected error, got none")
		}

		if err != nil && !errors.Is(err, s.listErr) {
			t.Errorf("unexpected error: %v", err)
		}

		if expenses == nil {
			t.Error("expected expenses to be not nil")
		}

		if len(expenses) != 0 {
			t.Errorf("expected expenses to be empty, got: %v", expenses)
		}
	})

	t.Run("successfully returns expenses", func(t *testing.T) {
		s := newMockStorage()

		service := ExpenseService{expenseStorage: s}
		expense1, _ := service.AddExpense("expense 1", 10)
		expense2, _ := service.AddExpense("expense 2", 20)

		expenses, err := service.ListExpenses()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if expenses == nil {
			t.Error("expected expenses to be not nil")
		}

		if len(expenses) != 2 {
			t.Errorf("expected expenses to be not empty, got: %v", expenses)
		}

		if expenses[0] != *expense1 {
			t.Errorf("expected expense 1, got: %v", expenses[0])
		}

		if expenses[1] != *expense2 {
			t.Errorf("expected expense 2, got: %v", expenses[1])
		}
	})
}

func TestExpenseService_ExpenseSummary(t *testing.T) {
	s := newMockStorage()

	service := ExpenseService{expenseStorage: s}
	service.AddExpense("expense 1", 10)
	service.AddExpense("expense 2", 20)

	total, err := service.ExpenseSummary()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if total != 30 {
		t.Errorf("expected total to be 30, got: %v", total)
	}

	if total != 30 {
		t.Errorf("expected total to be 30, got: %v", total)
	}
}
