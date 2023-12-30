package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/spf13/cobra"
)

func init() {
	// Added the balance sub command to the user command
	userCmd.AddCommand(balanceCmd)
}

type UserBalResponse struct {
	Token    string      `json:"token"`
	UserData UserBalData `json:"userData"`
}

type UserBalData struct {
	Balance UserBal `json:"balance"`
}

type UserBal struct {
	CurrencyCode string `json:"currencyCode"`
	Units        string `json:"units"`
	Nanos        int32  `json:"nanos"`
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
		// Check if user has already login
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
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Fetching user balance...")
		w.Start()
		client := http.Client{}
		req, err := http.NewRequest("GET", "https://okane-production.up.railway.app/api/user/"+userData.Email, nil)
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
			var respData UserBalResponse
			err = json.Unmarshal([]byte(jsonStr), &respData)
			if err != nil {
				log.Fatalf(err.Error())
			}
			// Gen Balance
			balUnitStr := strconv.Itoa(int(respData.UserData.Balance.Nanos))
			if len(balUnitStr) > 3 {
				if string(balUnitStr[0]) == "-" {
					balUnitStr = balUnitStr[1:4]
				} else {
					balUnitStr = balUnitStr[:3]
				}
			}
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Fetched Successfully!")
			fmt.Println()
			fmt.Println("User Balance =========")
			fmt.Println(respData.UserData.Balance.Units + "." + balUnitStr + " USD")
		}
	},
}
