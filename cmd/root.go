package cmd

import (
	"github.com/spf13/cobra"
	"github.com/toramanomer/expense-tracker/expense"
)

// commands holds the dependencies for all commands
type commands struct {
	service *expense.ExpenseService
}

// NewCommands creates a new Commands instance with the provided service
func NewCommands(service *expense.ExpenseService) *commands {
	return &commands{
		service: service,
	}
}

// RootCommand creates and returns the root command with all subcommands
func (c *commands) RootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "expense-tracker",
		Short: "A simple expense tracker CLI application to manage your finances.",
		Long:  "A command-line application to track your expenses, manage budgets, and generate reports.",
	}

	// Add all subcommands
	rootCmd.AddCommand(c.addCommand())
	rootCmd.AddCommand(c.deleteCommand())
	rootCmd.AddCommand(c.listCommand())
	rootCmd.AddCommand(c.summaryCommand())

	return rootCmd
}
