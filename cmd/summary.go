/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Display total expenses or monthly summary",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		m, _ := cmd.Flags().GetInt("month")
		month := time.Month(m)

		if m != 0 && (month < time.January || month > time.December) {
			fmt.Println("Invalid month. Please enter a valid month (1-12).")
			return
		}

		expenses, err := service.ListExpenses()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var totalExpenses int
		for _, expense := range expenses {
			if m == 0 || (time.Now().Year() == expense.Date.Year() && month == expense.Date.Month()) {
				totalExpenses += expense.Amount
			}
		}

		if m == 0 {
			fmt.Printf("Total expenses: $%d\n", totalExpenses)
		} else {
			fmt.Printf("Monthly summary for %s %d: $%d\n", month.String(), time.Now().Year(), totalExpenses)
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().IntP("month", "m", 0, "month for summary (1-12)")
}
