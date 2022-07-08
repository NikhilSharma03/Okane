package datastruct

import (
	"encoding/json"
)

// The User struct defines User Model
type User struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

// The NewUser takes id, name, email, password, balance as param
// It returns a new User struct
func NewUser(Id, name, email, password string, balance float64) *User {
	return &User{
		ID:       Id,
		Name:     name,
		Email:    email,
		Password: password,
		Balance:  balance,
	}
}

// ToJSON method is available on User struct which marshals it to json
func (user *User) ToJSON() (string, error) {
	uBytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(uBytes), nil
}

// Unmarshal method is available on User struct which unmarshal the provided data to the User struct
func (user *User) Unmarshal(data string) error {
	return json.Unmarshal([]byte(data), user)
}
