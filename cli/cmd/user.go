package cmd

import (
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func init() {
	// Add the user sub command to the root command
	rootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage user related queries",
	Long: `
The user sub command is used to do user related queries such as

✪ Get User Details
Example:
	okane user info

✪ Get User Balance
Example:
	okane user balance 
 
✪ Create User
Example:
	okane user new   

✪ Login User
Example:
 	okane user login  

✪ Logout User
Example:
	okane user logout   
 
✪ Update User
Example:
 	okane user update   
 
✪ Delete User
Example:
 	okane user delete   
 
`,
}

type LoginUserData struct {
	Token           string `yaml:"token"`
	Name            string `yaml:"name"`
	ID              string `yaml:"id"`
	Email           string `yaml:"email"`
	Balance         string `yaml:"balance"`
	IsAuthenticated bool   `yaml:"is_authenticated"`
}

var loginUserData LoginUserData

func (lu *LoginUserData) Login(token, name, id, email, password, balance string) error {
	lu.Token = token
	lu.Name = name
	lu.ID = id
	lu.Email = email
	lu.Balance = balance
	lu.IsAuthenticated = true

	y, err := yaml.Marshal(lu)
	if err != nil {
		return err
	}

	return ioutil.WriteFile("cli/cred.yaml", y, 0644)
}

func (l *LoginUserData) LogOut() error {
	return os.Remove("cli/cred.yaml")
}
