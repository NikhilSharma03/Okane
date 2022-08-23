package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/spf13/cobra"
)

func init() {
	// Added the info sub command to the user command
	userCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows user details",
	Long: `
The info command is used to see current login user details
 
Example:
	okane user info 
`,
	Run: func(cmd *cobra.Command, args []string) {
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Fetching user details...")
		w.Start()
		// Check if user has already login
		userData, err := loginUserData.GetData()
		if err != nil {
			log.Fatalf(err.Error())
		}
		w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Fetched Successfully!")
		fmt.Println()
		fmt.Println("User Details =========")
		fmt.Println()
		fmt.Println("UserID :", userData.ID)
		fmt.Println("Name :", userData.Name)
		fmt.Println("Email :", userData.Email)
		fmt.Println("Balance :", userData.Balance)
	},
}
