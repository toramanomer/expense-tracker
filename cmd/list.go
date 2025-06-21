/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

		for _, expense := range expenses {
			fmt.Printf("%d\t%s\t%s\t%s\t%d\n", expense.ID, expense.Date, expense.Category, expense.Description, expense.Amount)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
