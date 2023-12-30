package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "okane",
	Short: "A CLI application to manage expenses",
	Long: `	
Okane is an CLI application which helps to manage expenses.
	
It has many functionalities such as
 ✪ User Authentication
 ✪ Managing your expenses
 ✪ Credit to expenses 
 ✪ Debit from expenses
 ✪ Keep track of total amount present
 ✪ History of transactions   

Commands:
 ✪ user - Do user related queries
 ✪ expense - Do expense related queries
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
