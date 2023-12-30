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
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	// Added the history sub command to the expense command
	expenseCmd.AddCommand(expHistoryCmd)
}

type ExpenseHistoryResponse struct {
	Message      string    `json:"string"`
	ExpensesData []Expense `json:"expensesData"`
}

type Expense struct {
	ID          string     `json:"id"`
	UserID      string     `json:"userId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Amount      ExpenseAmt `json:"amount"`
	Type        string     `json:"type"`
}

type ExpenseAmt struct {
	CurrencyCode string `json:"currencyCode"`
	Units        string `json:"units"`
	Nanos        int    `json:"nanos"`
}

var expHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get transaction history",
	Long: `
The history command is used to get transaction history

Example:
	okane expense history
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if user is logged in
		userData, err := loginUserData.GetData()
		if err != nil {
			log.Fatalf(err.Error())
		}
		// Make API Request
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Fetching transactions...")
		w.Start()
		client := http.Client{}
		req, err := http.NewRequest(http.MethodGet, "https://okane-production.up.railway.app/api/expense/user/"+userData.ID, nil)
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
			var transactions ExpenseHistoryResponse
			err := json.Unmarshal([]byte(jsonStr), &transactions)
			if err != nil {
				w.PersistWith(spin.Spinner{Frames: []string{""}}, "something went wrong. Failed to fetch transactions")
				log.Fatalf("")
			}
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Fetched transactions successfully!")
			fmt.Println()
			var formatData [][]string
			for _, expense := range transactions.ExpensesData {
				units := expense.Amount.Units
				nanos := strconv.Itoa(expense.Amount.Nanos)
				if len(nanos) > 3 {
					if string(nanos[0]) == "-" {
						nanos = nanos[1:4]
						if string(expense.Amount.Units[0]) != "-" {
							units = "-" + units
						}
					} else {
						nanos = nanos[:3]
					}
				}
				expAmount := units + "." + nanos + " " + expense.Amount.CurrencyCode
				expField := []string{expense.ID, expense.Title, expense.Description, expAmount, expense.Type}
				formatData = append(formatData, expField)
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Title", "Description", "Amount", "Type"})
			table.SetRowLine(true)
			table.AppendBulk(formatData)
			table.Render()
		}
	},
}
