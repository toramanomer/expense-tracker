/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/toramanomer/expense-tracker/expense"
)

func printExpensesTable(expenses []expense.Expense) {
	if len(expenses) == 0 {
		fmt.Println("No expenses to display.")
		return
	}

	// Define column widths
	idWidth := 4
	amountWidth := 8
	categoryWidth := 12
	dateWidth := 12
	descWidth := 25

	// Calculate dynamic widths based on content
	for _, expense := range expenses {
		if len(expense.Category) > categoryWidth {
			categoryWidth = len(expense.Category)
		}
		if len(expense.Description) > descWidth {
			descWidth = len(expense.Description)
		}
	}

	// Print header
	fmt.Printf("┌─%s─┬─%s─┬─%s─┬─%s─┬─%s─┐\n",
		strings.Repeat("─", idWidth),
		strings.Repeat("─", amountWidth),
		strings.Repeat("─", categoryWidth),
		strings.Repeat("─", dateWidth),
		strings.Repeat("─", descWidth),
	)

	fmt.Printf("│ %-*s │ %-*s │ %-*s │ %-*s │ %-*s │\n",
		idWidth, "ID",
		amountWidth, "Amount",
		categoryWidth, "Category",
		dateWidth, "Date",
		descWidth, "Description",
	)

	fmt.Printf("├─%s─┼─%s─┼─%s─┼─%s─┼─%s─┤\n",
		strings.Repeat("─", idWidth),
		strings.Repeat("─", amountWidth),
		strings.Repeat("─", categoryWidth),
		strings.Repeat("─", dateWidth),
		strings.Repeat("─", descWidth),
	)

	// Print rows
	for _, expense := range expenses {
		// Truncate description if too long
		desc := expense.Description
		if len(desc) > descWidth {
			desc = desc[:descWidth-3] + "..."
		}

		fmt.Printf("│ %-*d │ $%-*d │ %-*s │ %-*s │ %-*s │\n",
			idWidth, expense.ID,
			amountWidth-1, expense.Amount, // -1 for the $ sign
			categoryWidth, expense.Category,
			dateWidth, expense.Date.Format("2006-01-02"),
			descWidth, desc,
		)
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Long:  "Display all expenses in a formatted table",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := service.ListExpenses()
		if err != nil {
			fmt.Println("Error listing expenses:", err)
			return
		}

		printExpensesTable(expenses)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
