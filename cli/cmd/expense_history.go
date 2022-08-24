package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/spf13/cobra"
)

func init() {
	// Added the history sub command to the expense command
	expenseCmd.AddCommand(expHistoryCmd)
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
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/api/expense/user/"+userData.ID, nil)
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
			w.PersistWith(spin.Spinner{Frames: []string{""}}, resErr.Message)
		} else {
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Fetched transactions successfully!")
			fmt.Println(jsonStr)
		}
	},
}
