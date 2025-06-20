/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Display total expenses or monthly summary",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		month, _ := cmd.Flags().GetInt("month")
		fmt.Println("Month:", month)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().IntP("month", "m", 0, "month for summary (1-12)")
}
