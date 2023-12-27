package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

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
	IsAuthenticated bool   `yaml:"is_authenticated"`
}

var loginUserData LoginUserData

func (lu *LoginUserData) Login(token, name, id, email, password string) error {
	lu.Token = token
	lu.Name = name
	lu.ID = id
	lu.Email = email
	lu.IsAuthenticated = true

	y, err := yaml.Marshal(lu)
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("something went wrong while login")
	}

	configDir := filepath.Join(homeDir, ".config", "okane")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return fmt.Errorf("something went wrong while login")
	}

	filePath := filepath.Join(configDir, "cred.yml")

	return os.WriteFile(filePath, y, 0644)
}

func (l *LoginUserData) LogOut() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("something went wrong while logout")
	}

	filePath := filepath.Join(homeDir, ".config", "okane", "cred.yml")

	return os.Remove(filePath)
}

func (l *LoginUserData) GetData() (*LoginUserData, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to read cred file ")
	}

	configDir := filepath.Join(homeDir, ".config", "okane")
	filePath := filepath.Join(configDir, "cred.yml")

	// Check if cred.yaml exists
	_, err = os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("please login")
	}

	viper.SetConfigName("cred")
	viper.AddConfigPath(configDir)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read cred file! please remove the cred file and login again",
		)
	}
	err = viper.Unmarshal(l)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to unmarshal cred file! please remove the cred file and login again",
		)
	}

	return l, nil
}
