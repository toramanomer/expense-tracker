package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toramanomer/expense-tracker/expense"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete an expense by ID",
	Args:    cobra.NoArgs,
	Example: "expense-tracker delete --id 2",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if err := expense.ValidateID(id); err != nil {
			fmt.Println(err)
			return
		}

		err := service.DeleteExpense(id)

		if err != nil {
			fmt.Printf("Error deleting expense with ID %d: %v\n", id, err)
			return
		}

		fmt.Printf("Expense with ID %d deleted successfully\n", id)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().Int("id", 0, "Expense ID to delete (required)")
	deleteCmd.MarkFlagRequired("id")
}
