package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/spf13/cobra"
)

func init() {
	// Added the login sub command to the user command
	userCmd.AddCommand(loginCmd)
}

type UserLoginResponse struct {
	Token    string        `json:"token"`
	UserData UserLoginData `json:"userData"`
}

type UserLoginData struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login as a user",
	Long: `
The login command is used to login as a user
 
Example:
	okane user login 
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Getting user email
		var email string
		emailPrompt := &survey.Input{
			Message: "Please type your email (email@okane.com) :",
		}
		survey.AskOne(emailPrompt, &email)
		// Getting user password
		var password string
		passPrompt := &survey.Password{
			Message: "Please type your password (min length: 6) :",
		}
		survey.AskOne(passPrompt, &password)
		// Validate data
		valEmail := strings.ToLower(strings.TrimSpace(email))
		valPassword := strings.TrimSpace(password)
		// Check for empty field
		if valEmail == "" || valPassword == "" {
			fmt.Println("Empty input found. Please type correct information.")
			os.Exit(1)
		}
		// Check for valid email
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@okane.com$`)
		if !emailRegex.MatchString(valEmail) {
			fmt.Println("Invalid email. (please use email@okane.com format)")
			os.Exit(1)
		}
		// Check for password length
		if len(valPassword) < 6 {
			fmt.Println("Password length should be atleast six.")
			os.Exit(1)
		}
		// Make API Request
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Login user...")
		w.Start()
		client := http.Client{}
		req, err := http.NewRequest("GET", "http://localhost:8000/api/user/"+valEmail, nil)
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
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf(err.Error())
		}
		jsonStr := string(body)

		if strings.Contains(jsonStr, "code") {
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
				log.Fatalf("Failed to login! %v", err.Error())
			}
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Login Successfully!")
		}
	},
}
