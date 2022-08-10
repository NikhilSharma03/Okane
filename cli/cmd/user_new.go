package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func init() {
	// Add the new sub command to the user command
	userCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new user",
	Long: `
The new command is used to create a new user
 
Example:
	okane user new 
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Getting user name
		var name string
		namePrompt := &survey.Input{
			Message: "Please type your name :",
		}
		survey.AskOne(namePrompt, &name)
		// Getting user email
		var email string
		emailPrompt := &survey.Input{
			Message: "Please type your email :",
		}
		survey.AskOne(emailPrompt, &email)
		// Getting user password
		var password string
		passPrompt := &survey.Password{
			Message: "Please type your password (min length: 6) :",
		}
		survey.AskOne(passPrompt, &password)

		// Validate data
		valName := strings.TrimSpace(name)
		valEmail := strings.TrimSpace(email)
		valPassword := strings.TrimSpace(password)
		// Check for empty field
		if valName == "" || valEmail == "" || valPassword == "" {
			fmt.Println("Empty input found. Please type correct information.")
			os.Exit(1)
		}
		// Check for password length
		if len(valPassword) < 6 {
			fmt.Println("Password length should be atleast six.")
			os.Exit(1)
		}
		// Check for valid email

		fmt.Println(valName, valEmail, valPassword)
	},
}
