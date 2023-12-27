package cmd

import (
	"bytes"
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

type ExpenseNewRequest struct {
	ExpenseData *ExpenseData `json:"expense_data"`
}

type ExpenseData struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Amount      *ExpenseDataAmt `json:"amount"`
	Type        int             `json:"type"`
}

type ExpenseDataAmt struct {
	Units string `json:"units"`
	Nanos string `json:"nanos"`
}

func init() {
	// Added the new sub command to the expense command
	expenseCmd.AddCommand(expNewCmd)
}

var expNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new expense",
	Long: `
The new command is used to create a new expense

Example:
	okane expense new
`,
	Run: func(_ *cobra.Command, _ []string) {
		// Check if user is logged in
		userData, err := loginUserData.GetData()
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Getting expense title
		var title string
		titlePrompt := &survey.Input{
			Message: "Please type expense title :",
		}
		err = survey.AskOne(titlePrompt, &title)
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Getting expense description
		var description string
		descriptionPrompt := &survey.Input{
			Message: "Please type expense description :",
		}
		err = survey.AskOne(descriptionPrompt, &description)
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Getting expense amount
		var amount string
		amountPrompt := &survey.Input{
			Message: "Please type expense amount (USD) :",
		}
		err = survey.AskOne(amountPrompt, &amount)
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Getting expense type
		expType := ""
		expTypePrompt := &survey.Select{
			Message: "Choose expense type:",
			Options: []string{"Credit", "Debit"},
		}
		err = survey.AskOne(expTypePrompt, &expType)
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Validate data
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
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Creating new expense...")
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
		req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/api/expense", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Grpc-metadata-token", userData.Token)
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
			if resErr.Message == "your token has been expired" {
				resErr.Message += "! Please login again."
			}
			w.PersistWith(spin.Spinner{Frames: []string{""}}, resErr.Message)
		} else {
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Expense Created Successfully!")
		}
	},
}
