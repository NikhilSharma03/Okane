package service

import (
	"fmt"
	"log"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/NikhilSharma03/Okane/server/internal/repository"
	"github.com/google/uuid"
)

// The ExpenseService interface defines methods to implement
type ExpenseService interface {
	CreateExpense(userID, title, description string, amount *datastruct.Amount, expenseType datastruct.EXPENSE_TYPE) (*datastruct.Expense, error)
}

// The expenseService struct take dao and logger (lg)
type expenseService struct {
	dao repository.DAO
	lg  *log.Logger
}

// The NewExpenseService takes dao and lg as parameter and returns ExpenseService interface implemented struct
func NewExpenseService(dao repository.DAO, lg *log.Logger) ExpenseService {
	return &expenseService{dao, lg}
}

func (es *expenseService) CreateExpense(userID, title, description string, amount *datastruct.Amount, expenseType datastruct.EXPENSE_TYPE) (*datastruct.Expense, error) {
	// Generate expenseID
	id := uuid.New().String()
	if id == "" {
		es.lg.Printf("Failed to generate id")
		return nil, fmt.Errorf("failed to generate expense ID")
	}
	// Create a new expense
	expenseData := datastruct.NewExpense(id, userID, title, description, expenseType)
	expenseData.Amount = amount
	// Store expense in DB
	return es.dao.NewExpenseCollection().CreateExpense(expenseData)
}
