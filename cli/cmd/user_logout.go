package cmd

import (
	"log"
	"os"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/spf13/cobra"
)

func init() {
	// Added the logout sub command to the user command
	userCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout user",
	Long: `
The logout command is used to logout current login user.
 
Example:
	okane user logout 
`,
	Run: func(cmd *cobra.Command, args []string) {
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Logout...")
		w.Start()
		// Check if user has already login
		userData, err := loginUserData.GetData()
		if err != nil {
			log.Fatalf(err.Error())
		}
		err = userData.LogOut()
		if err != nil {
			log.Fatalf("failed to log out! %v", err.Error())
		}
		w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Logout Successful!")
	},
}
