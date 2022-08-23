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
	// Added the balance sub command to the user command
	userCmd.AddCommand(balanceCmd)
}

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Shows user balance",
	Long: `
The balance command is used to see current login user balance
 
Example:
	okane user balance 
`,
	Run: func(cmd *cobra.Command, args []string) {
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Fetching user balance...")
		w.Start()
		// Check if user has already login
		userData, err := loginUserData.GetData()
		if err != nil {
			log.Fatalf(err.Error())
		}
		w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Fetched Successfully!")
		fmt.Println()
		fmt.Println("User Balance =========")
		fmt.Println()
		fmt.Println(userData.Balance)
	},
}
