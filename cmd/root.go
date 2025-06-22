package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/toramanomer/expense-tracker/expense"
)

var storage = expense.NewStorageFS("data")
var service = expense.NewExpenseService(storage)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "A simple expense tracker CLI application to manage your finances.",
	Long:  "A command-line application to track your expenses, manage budgets, and generate reports.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
