package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/spf13/cobra"
)

func init() {
	// Added the delete sub command to the expense command
	expenseCmd.AddCommand(expDeleteCmd)
	expDeleteCmd.Flags().StringP("id", "I", "", "The id flag is used to define expenses id")
}

var expDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete single transaction",
	Long: `
The delete command is used to delete single transaction
 
Example:
	okane expense delete --id/-I {expenseID}
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
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Deleting transaction...")
		w.Start()
		client := http.Client{}
		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/api/expense/"+expenseID, nil)
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
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Deleted transaction successfully!")
		}
	},
}
