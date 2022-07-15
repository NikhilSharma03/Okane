package datastruct

import (
	"encoding/json"
)

// The Balance struct define Balance model
type Balance struct {
	// The three-letter currency code
	CurrencyCode string `json:"currency_code"`
	// The whole units of the amount.
	Units int64 `json:"units"`
	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos int32 `json:"nanos"`
}

// The NewBalance takes currency code, units, nanos as param
// It returns a new Balance struct
func NewBalance(currencyCode string, units int64, nanos int32) *Balance {
	return &Balance{
		CurrencyCode: currencyCode,
		Units:        units,
		Nanos:        nanos,
	}
}

// The User struct defines User Model
type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Balance  *Balance `json:"balance"`
}

// The NewUser takes id, name, email, password, balance as param
// It returns a new User struct
func NewUser(Id, name, email, password string) *User {
	return &User{
		ID:       Id,
		Name:     name,
		Email:    email,
		Password: password,
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
