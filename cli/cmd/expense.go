package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// Add the expense sub command to the root command
	rootCmd.AddCommand(expenseCmd)
}

var expenseCmd = &cobra.Command{
	Use:   "expense",
	Short: "Manage expense related queries",
	Long: `
The expense sub command is used to do expense related queries such as

✪ Create Expense
Example:
	okane expense new   

✪ Get All Expenses (Transaction History)
Example:
	okane expense history

✪ Get Single Expense
Example:
	okane expense get --id/-I {expense ID}  
 
✪ Update Single Expense
Example:
 	okane expense update --id/-I {expense ID}   
 
✪ Delete Single Expense
Example:
 	okane expense delete --id/-I {expense ID}   
 
`,
}
