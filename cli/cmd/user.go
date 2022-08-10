package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// Add the user sub command to the root command
	rootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage user related queries",
	Long: `
The user sub command is used to do user related queries such as

✪ Get User Details
Example:
	okane user info

✪ Get User Balance
Example:
	okane user balance 
 
✪ Create User
Example:
	okane user new   

✪ Login User
Example:
 	okane user login  

✪ Logout User
Example:
	okane user logout   
 
✪ Update User
Example:
 	okane user update   
 
✪ Delete User
Example:
 	okane user delete   
 
`,
}
