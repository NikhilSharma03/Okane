package datastruct

import (
	"encoding/json"
)

// The Amount struct define Amount model
type Amount struct {
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

// The NewAmount takes currency code, units, nanos as param
// It returns a new Amount struct
func NewAmount(currencyCode string, units int64, nanos int32) *Amount {
	return &Amount{
		CurrencyCode: currencyCode,
		Units:        units,
		Nanos:        nanos,
	}
}

type EXPENSE_TYPE int32

const (
	CREDIT EXPENSE_TYPE = 0
	DEBIT  EXPENSE_TYPE = 1
)

// The Expense struct defines Expense Model
type Expense struct {
	Id          string       `json:"id"`
	UserId      string       `json:"user_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Amount      *Amount      `json:"amount"`
	Type        EXPENSE_TYPE `json:"type"`
}

// The NewExpense takes  as param
// It returns a new Expense struct
func NewExpense(id, userid, title, desc string, typ EXPENSE_TYPE) *Expense {
	return &Expense{
		Id:          id,
		UserId:      userid,
		Title:       title,
		Description: desc,
		Type:        typ,
	}
}

// ToJSON method is available on Expense struct which marshals it to json
func (expense *Expense) ToJSON() (string, error) {
	uBytes, err := json.Marshal(expense)
	if err != nil {
		return "", err
	}
	return string(uBytes), nil
}

// Unmarshal method is available on Expense struct which unmarshal the provided data to the Expense struct
func (expense *Expense) Unmarshal(data string) error {
	return json.Unmarshal([]byte(data), expense)
}
