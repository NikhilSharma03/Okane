package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/go-redis/redis/v8"
)

// The ExpenseCollection defines the methods a struct need to have
type ExpenseCollection interface {
	CreateExpense(expense_data *datastruct.Expense) (*datastruct.Expense, error)
	GetExpenses(userID string) ([]*datastruct.Expense, error)
	GetExpenseByID(expenseID string) (*datastruct.Expense, error)
	DeleteExpenseByID(userID, expenseID string) error
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
	expenseData, err := json.Marshal(expense_data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal expense data")
	}
	// Save the expense in the DB
	_, err = DB.HSet(context.Background(), EXPENSE_COLLECTION, EXPENSE+expense_data.Id, expenseData).Result()
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

func (*expenseCollection) GetExpenses(userID string) ([]*datastruct.Expense, error) {
	// The result userExpenses
	userExpenses := []*datastruct.Expense{}
	// Get all the expense ID's of expenses created by provided userID
	expenseIDs, err := DB.ZRevRange(context.Background(), USER_EXPENSE+userID, 0, -1).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no expense ids exists with provided userID")
		}
		return nil, fmt.Errorf("failed to fetch expense ids")
	}
	for _, expenseID := range expenseIDs {
		// Find Expense using the expenseID
		expString, err := DB.HGet(context.Background(), EXPENSE_COLLECTION, EXPENSE+expenseID).Result()
		if err != nil {
			if err == redis.Nil {
				return nil, fmt.Errorf("no expense exists with provided expenseID")
			}
			return nil, fmt.Errorf("failed to fetch expense with provided id")
		}
		// Unmarshal to expense
		var expense datastruct.Expense
		err = expense.Unmarshal(expString)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal expense data")
		}
		// Store expense in userExpenses
		userExpenses = append(userExpenses, &expense)
	}
	return userExpenses, nil
}

func (*expenseCollection) GetExpenseByID(expenseID string) (*datastruct.Expense, error) {
	// Fetch expense with provided ID
	expString, err := DB.HGet(context.Background(), EXPENSE_COLLECTION, EXPENSE+expenseID).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no expense exists with provided expenseID")
		}
		return nil, fmt.Errorf("failed to fetch expense with provided id")
	}
	// Unmarshal to expense
	var expense datastruct.Expense
	err = expense.Unmarshal(expString)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal expense data")
	}
	return &expense, nil
}

func (*expenseCollection) DeleteExpenseByID(userID, expenseID string) error {
	// Remove expense from DB
	_, err := DB.HDel(context.Background(), EXPENSE_COLLECTION, EXPENSE+expenseID).Result()
	if err != nil {
		return fmt.Errorf("failed to remove expense")
	}
	// Remove the expense ID from user_expense
	_, err = DB.ZRem(context.Background(), USER_EXPENSE+userID, expenseID).Result()
	if err != nil {
		return fmt.Errorf("failed to remove expense from user_expense")
	}
	return nil
}
