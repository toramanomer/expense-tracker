package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toramanomer/expense-tracker/expense"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new expense",
	Long:    "Add a new expense with description, amount (in dollars), and category",
	Example: "expense-tracker add --category \"Food\" --description \"Lunch\" --amount 20",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		if c, err := expense.ParseCategory(category); err != nil {
			fmt.Println(err)
			return
		} else {
			category = c
		}

		description, _ := cmd.Flags().GetString("description")
		if d, err := expense.ParseDescription(description); err != nil {
			fmt.Println(err)
			return
		} else {
			description = d
		}

		amount, _ := cmd.Flags().GetInt("amount")
		if err := expense.ValidateAmount(amount); err != nil {
			fmt.Println(err)
			return
		}

		expense, err := service.AddExpense(category, description, amount)
		if err != nil {
			fmt.Println("Error adding expense:", err)
			return
		}

		fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("category", "c", "", "Expense category (required)")
	addCmd.Flags().StringP("description", "d", "", "Expense description (required)")
	addCmd.Flags().IntP("amount", "a", 0, "Expense amount (required)")

	addCmd.MarkFlagRequired("category")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
