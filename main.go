package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"

	"github.com/toramanomer/expense-tracker/cmd"
	"github.com/toramanomer/expense-tracker/expense"
)

func main() {
	var (
		storage  = expense.NewStorageFS("data")
		service  = expense.NewExpenseService(storage)
		commands = cmd.NewCommands(service)
	)

	if err := fang.Execute(context.Background(), commands.RootCommand()); err != nil {
		os.Exit(1)
	}
}
