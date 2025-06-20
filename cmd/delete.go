package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete an expense by ID",
	Args:    cobra.NoArgs,
	Example: "expense-tracker delete --id 2",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		fmt.Printf("Deleting expense with ID: %d\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntP("id", "i", 0, "Expense ID to delete (required)")
	deleteCmd.MarkFlagRequired("id")
}
