package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/spf13/cobra"
)

func init() {
	// Added the delete sub command to the user command
	userCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user",
	Long: `
The delete command is used to delete the current login user

Example:
	okane user delete
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if user is login
		userData, err := loginUserData.GetData()
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Getting user password
		var password string
		passPrompt := &survey.Password{
			Message: "Please type your password (for authentication) :",
		}
		err = survey.AskOne(passPrompt, &password)
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Validate data
		valPassword := strings.TrimSpace(password)
		// Check for empty field
		if valPassword == "" {
			fmt.Println("Empty input found. Please type correct information.")
			os.Exit(1)
		}
		// Check for password length
		if len(valPassword) < 6 {
			fmt.Println("Password length should be atleast six.")
			os.Exit(1)
		}
		// Make API Request
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Deleting user...")
		w.Start()
		client := http.Client{}
		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/api/user/"+userData.Email, nil)
		if err != nil {
			log.Fatalf(err.Error())
		}
		req.Header = http.Header{
			"Grpc-metadata-cred": {valPassword},
		}
		res, err := client.Do(req)
		if err != nil {
			log.Fatalf(err.Error())
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf(err.Error())
		}
		jsonStr := string(body)

		if res.StatusCode == http.StatusInternalServerError {
			var resErr ResponseError
			err = json.Unmarshal([]byte(jsonStr), &resErr)
			if err != nil {
				log.Fatalf(err.Error())
			}
			w.PersistWith(spin.Spinner{Frames: []string{""}}, resErr.Message)
		} else {
			err = userData.LogOut()
			if err != nil {
				log.Fatalf("failed to remove cred file. please delete manually")
			}
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Deleted Successfully!")
		}
	},
}
