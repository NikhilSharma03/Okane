package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/go-redis/redis/v8"
)

// The ExpenseCollection defines the methods a struct need to have
type ExpenseCollection interface {
	CreateExpense(expense_data *datastruct.Expense) (*datastruct.Expense, error)
	GetExpense(userID string) ([]*datastruct.Expense, error)
	GetExpenseByID(expenseID string) (*datastruct.Expense, error)
	UpdateExpenseByID(expense_data *datastruct.Expense) (*datastruct.Expense, error)
	DeleteExpenseByID(expenseID string) (*datastruct.Expense, error)
}

// The expenseCollection struct implements method as defined in ExpenseCollection interface
type expenseCollection struct{}

// Expense collection const
const (
	EXPENSE_COLLECTION = "expense"
	EXPENSE            = "expense:"
	USER_EXPENSE       = "user_expense:"
)

func (*expenseCollection) CreateExpense(expense_data *datastruct.Expense) (*datastruct.Expense, error) {
	// Save the expense in the DB
	_, err := DB.HSet(context.Background(), EXPENSE_COLLECTION, EXPENSE+expense_data.Id, expense_data).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to create a new expense")
	}
	// Save the expense in user_expense collection in the DB
	_, err = DB.ZAdd(context.Background(), USER_EXPENSE+expense_data.UserId, &redis.Z{Score: float64(time.Now().Unix()), Member: expense_data.Id}).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to store expense id in user expense collection")
	}
	return expense_data, nil
}

func (*expenseCollection) GetExpense(userID string) ([]*datastruct.Expense, error)
func (*expenseCollection) GetExpenseByID(expenseID string) (*datastruct.Expense, error)
func (*expenseCollection) UpdateExpenseByID(expense_data *datastruct.Expense) (*datastruct.Expense, error)
func (*expenseCollection) DeleteExpenseByID(expenseID string) (*datastruct.Expense, error)
