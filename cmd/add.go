package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toramanomer/expense-tracker/expense"
)

// addCommand creates the add command
func (c *commands) addCommand() *cobra.Command {
	addCmd := &cobra.Command{
		Use:     "add",
		Short:   "Add a new expense",
		Long:    "Add a new expense with description, amount (in dollars), and category",
		Example: "expense-tracker add --category \"Food\" --description \"Lunch\" --amount 20",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			category, _ := cmd.Flags().GetString("category")
			if parsed, err := expense.ParseCategory(category); err != nil {
				fmt.Println(err)
				return
			} else {
				category = parsed
			}

			description, _ := cmd.Flags().GetString("description")
			if parsed, err := expense.ParseDescription(description); err != nil {
				fmt.Println(err)
				return
			} else {
				description = parsed
			}

			amount, _ := cmd.Flags().GetInt("amount")
			if err := expense.ValidateAmount(amount); err != nil {
				fmt.Println(err)
				return
			}

			expense, err := c.service.AddExpense(category, description, amount)
			if err != nil {
				fmt.Println("Error adding expense:", err)
				return
			}

			fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)
		},
	}

	addCmd.Flags().StringP("category", "c", "", "Expense category (required)")
	addCmd.Flags().StringP("description", "d", "", "Expense description (required)")
	addCmd.Flags().IntP("amount", "a", 0, "Expense amount (required)")

	addCmd.MarkFlagRequired("category")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")

	return addCmd
}
