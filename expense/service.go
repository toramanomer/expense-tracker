package expense

import "time"

type ExpenseService struct {
	expenseStorage ExpenseStorage
}

func NewExpenseService(expenseStorage ExpenseStorage) *ExpenseService {
	return &ExpenseService{
		expenseStorage: expenseStorage,
	}
}

// AddExpense adds a new expense with category, description and amount for today
func (s *ExpenseService) AddExpense(category, description string, amount int) (*Expense, error) {
	id, err := s.expenseStorage.GenerateID()
	if err != nil {
		return nil, err
	}

	expense := Expense{
		ID:          id,
		Date:        time.Now(),
		Category:    category,
		Description: description,
		Amount:      amount,
	}

	if err := s.expenseStorage.Add(expense); err != nil {
		return nil, err
	}

	return &expense, nil
}

// DeleteExpense deletes an expense by its ID
func (s *ExpenseService) DeleteExpense(id int) error {
	return s.expenseStorage.Delete(id)
}

// ListExpenses lists all expenses
func (s *ExpenseService) ListExpenses() ([]Expense, error) {
	expenses, err := s.expenseStorage.List()
	if err != nil {
		return []Expense{}, err
	}
	return expenses, nil
}

// ExpenseSummary calculates the total amount of all expenses
func (s *ExpenseService) ExpenseSummary() (int, error) {
	expenses, err := s.ListExpenses()
	if err != nil {
		return 0, err
	}

	var total int
	for _, expense := range expenses {
		total += expense.Amount
	}

	return total, nil
}
