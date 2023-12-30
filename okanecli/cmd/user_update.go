package cmd

import (
	"bytes"
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
	// Added the update sub command to the user command
	userCmd.AddCommand(updateCmd)
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update user details",
	Long: `
The update command is used to update current login user details

Example:
	okane user update
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if user is login
		userData, err := loginUserData.GetData()
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Get updated values
		// Getting user email
		var name string
		namePrompt := &survey.Input{
			Message: "Please type your name (to update) :",
		}
		err = survey.AskOne(namePrompt, &name)
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
		valName := strings.TrimSpace(name)
		valPassword := strings.TrimSpace(password)
		// Check for empty field
		if valName == "" || valPassword == "" {
			fmt.Println("Empty input found. Please type correct information.")
			os.Exit(1)
		}
		// Check for password length
		if len(valPassword) < 6 {
			fmt.Println("Password length should be atleast six.")
			os.Exit(1)
		}
		// Make API Request
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Updating user...")
		w.Start()
		reqBody := &UpdateUserRequest{
			Email:    userData.Email,
			Name:     valName,
			Password: valPassword,
		}
		body, err := json.Marshal(reqBody)
		if err != nil {
			log.Fatalf("Failed to marshal user request body")
		}
		client := http.Client{}
		req, err := http.NewRequest(http.MethodPatch, "https://okane-production.up.railway.app/api/user/"+userData.Email, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Fatalf(err.Error())
		}
		res, err := client.Do(req)
		if err != nil {
			log.Fatalf(err.Error())
		}
		body, err = io.ReadAll(res.Body)
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
			var respData UserLoginResponse
			err = json.Unmarshal([]byte(jsonStr), &respData)
			if err != nil {
				log.Fatalf(err.Error())
			}
			err = loginUserData.Login(respData.Token, respData.UserData.Name, respData.UserData.ID, respData.UserData.Email, valPassword)
			if err != nil {
				log.Fatalf("Failed to update new user detail! please manually delete cred file and login again")
			}
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Updated User Details Successfully!")
		}
	},
}
