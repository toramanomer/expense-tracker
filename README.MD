# Expense Tracker

A command-line application to track your expenses, manage budgets, and generate reports.

## Available Commands

Usage:
  `expense-tracker [command]`

Available Commands:
  `add`         Add a new expense
  `completion`  Generate the autocompletion script for the specified shell
  `delete`      Delete an expense by ID
  `help`        Help about any command
  `list`        List all expenses
  `summary`     Display total expenses or monthly summary

Flags:
  `-h`, `--help`   help for expense-tracker

Use `expense-tracker [command] --help` for more information about a command.

### Adding Expense

`add` command is used to add a new expand. It requires passing the expense amount, expense description, and expense category. The success message includes the ID of the newly added expense.

The following command adds a new expense:

```sh
expense-tracker add --category "Food" --description "Lunch" --amount 20
```

### Listing Expenses

`list` command is used to list all the expenses. It will output the ID, amount, description, category and date of each expense.

```sh
expense-tracker list
```

### Deleting Expense

`delete` command is used to delete an expense by its ID. The following command will delete an expense with ID 3:

```sh
expense-tracker delete 3
```


### Expense Summary

`summary` command is used to display total expenses for a given month in the current year or all expenses.

The following command displays the total expense:

```sh
expense-tracker summary
```


The following command displays the total expense for January this year. The month is specified by its number (1 for January, 2 for February, etc.).

```sh
expense-tracker summary 1
```
