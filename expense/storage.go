package expense

// ExpenseStorage interface defines the methods for managing expenses.
type ExpenseStorage interface {
	GenerateID() (int, error)  // Generates a unique ID for a new expense.
	Add(expense Expense) error // Adds a new expense to the storage.
	Delete(id int) error       // Deletes an expense from the storage.
	List() ([]Expense, error)  // Lists all expenses in the storage.
}
