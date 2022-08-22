package cmd

import (
	"bytes"
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

type UserNewRequest struct {
	UserData *UserData `json:"user_data"`
}
type UserData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserNewRespErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func init() {
	// Added the new sub command to the user command
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
		valName := strings.TrimSpace(name)
		valEmail := strings.ToLower(strings.TrimSpace(email))
		valPassword := strings.TrimSpace(password)
		// Check for empty field
		if valName == "" || valEmail == "" || valPassword == "" {
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
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Creating new user...")
		w.Start()
		reqBody := &UserNewRequest{
			UserData: &UserData{
				Name:     valName,
				Email:    valEmail,
				Password: valPassword,
			},
		}
		body, err := json.Marshal(reqBody)
		if err != nil {
			log.Fatalf("Failed to marshal user request body")
		}
		resp, err := http.Post("http://localhost:8000/api/user", "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf(err.Error())
		}
		jsonStr := string(body)
		if strings.Contains(jsonStr, "code") {
			var resErr UserNewRespErr
			err = json.Unmarshal([]byte(jsonStr), &resErr)
			if err != nil {
				log.Fatalf(err.Error())
			}
			w.PersistWith(spin.Spinner{Frames: []string{""}}, " !"+resErr.Message)
		} else {
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " User Created Successfully!")
		}
	},
}
