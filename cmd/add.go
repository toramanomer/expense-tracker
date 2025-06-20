package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new expense",
	Long:    "Add a new expense with description, amount (in dollars), and category",
	Example: "expense-tracker add --description \"Lunch\" --amount 20",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		description, _ := cmd.Flags().GetString("description")
		amount, _ := cmd.Flags().GetInt("amount")
		fmt.Println("description:", description)
		fmt.Println("amount:", amount)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("description", "d", "", "Expense description (required)")
	addCmd.Flags().IntP("amount", "a", 0, "Expense amount (required)")

	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
