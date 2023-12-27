package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// Added the update sub command to the expense command
	expenseCmd.AddCommand(expUpdateCmd)
	expUpdateCmd.Flags().StringP("id", "I", "", "The id flag is used to define expenses id")
}

type ExpenseUpdateResponse struct {
	Message      string  `json:"string"`
	ExpensesData Expense `json:"expensesData"`
}

var expUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update single transaction details",
	Long: `
The update command is used to update single transaction details

Example:
	okane expense update --id/-I {expenseID}
`,
	Run: func(cmd *cobra.Command, _ []string) {
		// Getting expense ID
		expenseID, err := cmd.Flags().GetString("id")
		if err != nil || expenseID == "" {
			log.Fatalf("please enter correct ID")
		}
		// Check if user is logged in
		userData, err := loginUserData.GetData()
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Make API Request
		client := http.Client{}
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/api/expense/"+expenseID, nil)
		req.Header.Set("Grpc-metadata-token", userData.Token)
		if err != nil {
			log.Fatalf(err.Error())
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
		if res.StatusCode == http.StatusInternalServerError {
			var resErr ResponseError
			err = json.Unmarshal([]byte(jsonStr), &resErr)
			if err != nil {
				log.Fatalf(err.Error())
			}
			if resErr.Message == "your token has been expired" {
				resErr.Message += "! Please login again."
			}
			log.Fatalf(resErr.Message)
		} else {
			var transaction ExpenseGetResponse
			err := json.Unmarshal([]byte(jsonStr), &transaction)
			if err != nil {
				log.Fatalf("something went wrong.")
			}
			units := transaction.ExpensesData.Amount.Units
			nanos := strconv.Itoa(transaction.ExpensesData.Amount.Nanos)
			if len(nanos) > 3 {
				if string(nanos[0]) == "-" {
					nanos = nanos[1:4]
					if string(transaction.ExpensesData.Amount.Units[0]) != "-" {
						units = "-" + units
					}
				} else {
					nanos = nanos[:3]
				}
			}
			// Get old values
			expAmount := units + "." + nanos + " " + transaction.ExpensesData.Amount.CurrencyCode
			Title := transaction.ExpensesData.Title
			Description := transaction.ExpensesData.Description
			Amount := expAmount
			Type := transaction.ExpensesData.Type

			// Get Updated Values
			// Getting expense title
			var title string
			titlePrompt := &survey.Input{
				Message: "Please type expense title (current : " + Title + ") :",
			}
			err = survey.AskOne(titlePrompt, &title)
			if err != nil {
				log.Fatalf(err.Error())
			}
			// Getting expense description
			var description string
			descriptionPrompt := &survey.Input{
				Message: "Please type expense description (current : " + Description + ") :",
			}
			err = survey.AskOne(descriptionPrompt, &description)
			if err != nil {
				log.Fatalf(err.Error())
			}
			// Getting expense amount
			var amount string
			amountPrompt := &survey.Input{
				Message: "Please type expense amount (USD) (current : " + Amount + ") :",
			}
			err = survey.AskOne(amountPrompt, &amount)
			if err != nil {
				log.Fatalf(err.Error())
			}
			// Getting expense type
			expType := ""
			expTypePrompt := &survey.Select{
				Message: "Choose expense type (current : " + Type + ") :",
				Options: []string{"Credit", "Debit"},
			}
			err = survey.AskOne(expTypePrompt, &expType)
			if err != nil {
				log.Fatalf(err.Error())
			}
			valTitle := strings.TrimSpace(title)
			valDescription := strings.TrimSpace(description)
			valAmount := strings.TrimSpace(amount)
			valExpType := strings.TrimSpace(expType)
			// Check for empty field
			if valTitle == "" || valDescription == "" || valAmount == "" || valExpType == "" {
				fmt.Println("Empty input found. Please type correct information.")
				os.Exit(1)
			}
			// Generating amount
			expAmtUnits := ""
			expAmtNanos := ""
			expSplit := strings.Split(valAmount, ".")
			if len(expSplit) > 2 {
				fmt.Println("Invalid Amount entered.")
				os.Exit(1)
			}
			if len(expSplit) == 1 {
				expAmtUnits = expSplit[0]
				expAmtNanos = "0"
			} else if len(expSplit) == 2 {
				expAmtUnits = expSplit[0]
				expAmtNanos = expSplit[1]
				expAmtNanosLen := len(expAmtNanos)
				for expAmtNanosLen < 9 {
					expAmtNanos += "0"
					expAmtNanosLen++
				}
				if strings.Contains(expAmtUnits, "-") {
					temp := "-" + expAmtNanos
					expAmtNanos = temp
				}
				eAUInt, err := strconv.Atoi(expAmtUnits)
				if err != nil {
					fmt.Println("Invalid Amount entered.")
					os.Exit(1)
				}
				eANInt, err := strconv.Atoi(expAmtNanos)
				if err != nil {
					fmt.Println("Invalid Amount entered.")
					os.Exit(1)
				}
				expAmtUnits = strconv.Itoa(eAUInt)
				expAmtNanos = strconv.Itoa(eANInt)
			}
			// Generating correct type
			var expenseType int
			if valExpType == "Credit" {
				expenseType = 0
			} else if valExpType == "Debit" {
				expenseType = 1
			}
			// Make API Request
			w := wow.New(os.Stdout, spin.Get(spin.Dots), " Updating expense...")
			w.Start()
			reqBody := &ExpenseNewRequest{
				ExpenseData: &ExpenseData{
					Title:       valTitle,
					Description: valDescription,
					Type:        expenseType,
					Amount: &ExpenseDataAmt{
						Units: expAmtUnits,
						Nanos: expAmtNanos,
					},
				},
			}
			body, err := json.Marshal(reqBody)
			if err != nil {
				log.Fatalf("Failed to marshal expense request body")
			}
			client := http.Client{}
			req, err := http.NewRequest(http.MethodPatch, "http://localhost:8000/api/expense/"+expenseID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Grpc-metadata-token", userData.Token)
			if err != nil {
				log.Fatalf(err.Error())
			}
			res, err := client.Do(req)
			if err != nil {
				log.Fatalf(err.Error())
			}
			body, err = ioutil.ReadAll(res.Body)
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
				if resErr.Message == "your token has been expired" {
					resErr.Message += "! Please login again."
				}
				w.PersistWith(spin.Spinner{Frames: []string{""}}, resErr.Message)
			} else {
				w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Expense Updated Successfully!")
			}
		}
	},
}
