package cmd

import "github.com/spf13/cobra"

func init() {
	// Add the user sub command to the root command
	rootCmd.AddCommand(userCmd)
	// Add flags to user sub command
	userCmd.Flags().StringP("info", "I", "", "The info flag is used to check user info")
	userCmd.Flags().StringP("balance", "B", "", "The balance flag is used to check user balance")
	userCmd.Flags().StringP("new", "N", "", "The new flag is used to create a new user")
	userCmd.Flags().StringP("login", "L", "", "The login flag is used to login user")
	userCmd.Flags().StringP("logout", "S", "", "The logout flag is used to logout user")
	userCmd.Flags().StringP("update", "U", "", "The update flag is used to update user")
	userCmd.Flags().StringP("delete", "D", "", "The delete flag is used to delete user")
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage user related queries",
	Long: `
 The user sub command is used to do user related queries such as

 ✪ Get User Details
 Example:
	okane user --info  // Can use either --info or -I flag

 ✪ Get User Balance
 Example:
	okane user --balance  // Can use either --balance or -B flag
 
 ✪ Create User
 Example:
	okane user --new   // Can use either --new or -N flag

 ✪ Login User
 Example:
 	okane user --login   // Can use either --login or -L flag

 ✪ Logout User
 Example:
	okane user --logout   // Can use either --logout or -S flag
 
 ✪ Update User
 Example:
 	okane user --update   // Can use either --update or -U flag
 
 ✪ Delete User
 Example:
 	okane user --delete   // Can use either --delete or -D flag
 
`,
}
