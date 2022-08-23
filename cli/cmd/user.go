package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	var filePath string = "cli/cred.yaml"
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("something went wrong while login")
	}
	pathLen := len(path)
	lastPath := string(path[pathLen-3]) + string(path[pathLen-2]) + string(path[pathLen-1])
	if lastPath == "cli" {
		filePath = "cred.yaml"
	} else if lastPath == "ane" {
		filePath = "cli/cred.yaml"
	} else {
		return fmt.Errorf("please execute cli app from root dir or cli dir")
	}

	return ioutil.WriteFile(filePath, y, 0644)
}

func (l *LoginUserData) LogOut() error {
	var filePath string = "cli/cred.yaml"
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("something went wrong while login")
	}
	pathLen := len(path)
	lastPath := string(path[pathLen-3]) + string(path[pathLen-2]) + string(path[pathLen-1])
	if lastPath == "cli" {
		filePath = "cred.yaml"
	} else if lastPath == "ane" {
		filePath = "cli/cred.yaml"
	} else {
		return fmt.Errorf("please execute cli app from root dir or cli dir")
	}
	return os.Remove(filePath)
}

func (l *LoginUserData) GetData() (*LoginUserData, error) {
	var filePath string = "cli/cred.yaml"
	var confPath string = "./cli"
	path, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("something went wrong while login")
	}
	pathLen := len(path)
	lastPath := string(path[pathLen-3]) + string(path[pathLen-2]) + string(path[pathLen-1])
	if lastPath == "cli" {
		filePath = "cred.yaml"
		confPath = "."
	} else if lastPath == "ane" {
		filePath = "cli/cred.yaml"
	} else {
		return nil, fmt.Errorf("please execute cli app from root dir or cli dir")
	}
	// Check if cred.yaml exists
	_, err = os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("please login")
	}
	// Get values from cred.yaml
	viper.SetConfigName("cred")
	viper.AddConfigPath(confPath)
	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read cred file! please remove the cred file and login again")
	}
	err = viper.Unmarshal(l)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cred file! please remove the cred file and login again")
	}
	return l, nil
}
