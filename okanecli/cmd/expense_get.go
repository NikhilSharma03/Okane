package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/spf13/cobra"
)

func init() {
	// Added the get sub command to the expense command
	expenseCmd.AddCommand(expGetCmd)
	expGetCmd.Flags().StringP("id", "I", "", "The id flag is used to define expenses id")
}

type ExpenseGetResponse struct {
	Message      string  `json:"string"`
	ExpensesData Expense `json:"expensesData"`
}

var expGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get single transaction details",
	Long: `
The get command is used to get single transaction details

Example:
	okane expense get --id/-I {expenseID}
`,
	Run: func(cmd *cobra.Command, args []string) {
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
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Fetching transaction details...")
		w.Start()
		client := http.Client{}
		req, err := http.NewRequest(http.MethodGet, "https://okane-production.up.railway.app/api/expense/"+expenseID, nil)
		req.Header.Set("Grpc-metadata-token", userData.Token)
		if err != nil {
			log.Fatalf(err.Error())
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
			if resErr.Message == "your token has been expired" {
				resErr.Message += "! Please login again."
			}
			w.PersistWith(spin.Spinner{Frames: []string{""}}, resErr.Message)
		} else {
			var transaction ExpenseGetResponse
			err := json.Unmarshal([]byte(jsonStr), &transaction)
			if err != nil {
				w.PersistWith(spin.Spinner{Frames: []string{""}}, "something went wrong. Failed to fetch transactions")
				log.Fatalf("")
			}
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Fetched transaction successfully!")
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
			expAmount := units + "." + nanos + " " + transaction.ExpensesData.Amount.CurrencyCode
			fmt.Println()
			fmt.Println("ID :", transaction.ExpensesData.ID)
			fmt.Println("Title :", transaction.ExpensesData.Title)
			fmt.Println("Description :", transaction.ExpensesData.Description)
			fmt.Println("Amount :", expAmount)
			fmt.Println("Type :", transaction.ExpensesData.Type)
			fmt.Println()
		}
	},
}
