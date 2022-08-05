package service

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/NikhilSharma03/Okane/server/internal/repository"
	"github.com/NikhilSharma03/Okane/server/internal/utils"
	"github.com/google/uuid"
)

// The ExpenseService interface defines methods to implement
type ExpenseService interface {
	CreateExpense(userID, email, title, description string, amount *datastruct.Amount, expenseType datastruct.EXPENSE_TYPE) (*datastruct.Expense, error)
	GetExpenses(userID string) ([]*datastruct.Expense, error)
	GetExpenseByID(expenseID string) (*datastruct.Expense, error)
	UpdateExpenseByID(oldExpense *datastruct.Expense, expenseID, userID, email, title, description string, amount *datastruct.Amount, expenseType datastruct.EXPENSE_TYPE) (*datastruct.Expense, error)
	DeleteExpenseByID(userID, email, expenseID string) error
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

func (es *expenseService) CreateExpense(userID, email, title, description string, amount *datastruct.Amount, expenseType datastruct.EXPENSE_TYPE) (*datastruct.Expense, error) {
	es.lg.Println("CreateExpense Called...")
	// Generate expenseID
	id := uuid.New().String()
	if id == "" {
		es.lg.Printf("Failed to generate id")
		return nil, fmt.Errorf("failed to generate expense ID")
	}
	// Create a new expense
	expenseData := datastruct.NewExpense(id, userID, title, description, expenseType)
	expenseData.Amount = amount

	// Update user balance
	err := es.dao.NewUserCollection().UpdateUserBalance(email, expenseData.Amount.Units, expenseData.Amount.Nanos, expenseData.Type)
	if err != nil {
		es.lg.Printf(err.Error())
		return nil, err
	}

	// Store expense in DB
	return es.dao.NewExpenseCollection().CreateExpense(expenseData)
}

func (es *expenseService) GetExpenses(userID string) ([]*datastruct.Expense, error) {
	es.lg.Println("GetExpenses Called...")
	return es.dao.NewExpenseCollection().GetExpenses(userID)
}

func (es *expenseService) GetExpenseByID(expenseID string) (*datastruct.Expense, error) {
	es.lg.Println("GetExpenseByID Called...")
	return es.dao.NewExpenseCollection().GetExpenseByID(expenseID)
}

func (es *expenseService) UpdateExpenseByID(oldExpense *datastruct.Expense, expenseID, userID, email, title, description string, amount *datastruct.Amount, expenseType datastruct.EXPENSE_TYPE) (*datastruct.Expense, error) {
	es.lg.Println("UpdateExpenseByID Called...")
	// Create updated expense
	expenseData := datastruct.NewExpense(expenseID, userID, title, description, expenseType)
	expenseData.Amount = amount

	// Calculate difference after update
	var finalAmountUnits int64
	var finalAmountNanos int64
	var finalAmountType datastruct.EXPENSE_TYPE

	if oldExpense.Type == expenseData.Type {
		// If there is no change in type
		oldExpenseUnits := oldExpense.Amount.Units
		oldExpenseNanos := oldExpense.Amount.Nanos
		newExpenseUnits := expenseData.Amount.Units
		newExpenseNanos := expenseData.Amount.Nanos
		if oldExpense.Type == datastruct.CREDIT {
			fAU, fAN, result, err := utils.Calculate(newExpenseUnits, oldExpenseUnits, newExpenseNanos, oldExpenseNanos, "sub")
			if err != nil {
				return nil, err
			}
			finalAmountType = datastruct.CREDIT
			if strings.Contains(result, "-") {
				finalAmountType = datastruct.DEBIT
				finalAmountUnits = int64(math.Abs(float64(fAU)))
				finalAmountNanos = int64(math.Abs(float64(fAN)))
			}
		} else if oldExpense.Type == datastruct.DEBIT {
			fAU, fAN, _, err := utils.Calculate(newExpenseUnits, oldExpenseUnits, newExpenseNanos, oldExpenseNanos, "sub")
			if err != nil {
				return nil, err
			}
			finalAmountType = datastruct.DEBIT
			finalAmountUnits = fAU
			finalAmountNanos = fAN
		}
	} else if oldExpense.Type == datastruct.CREDIT && expenseData.Type == datastruct.DEBIT {
		// If the old expense was Credit but was updated to debit
		finalAmountUnits = int64(oldExpense.Amount.Units + expenseData.Amount.Units)
		finalAmountNanos = int64(oldExpense.Amount.Nanos + expenseData.Amount.Nanos)
		if finalAmountNanos > 999999999 {
			finalAmountUnits += 1
			nanoString := strconv.Itoa(int(finalAmountNanos))
			newNanoString := nanoString[1:]
			newNanoInt, _ := strconv.Atoi(newNanoString)
			finalAmountNanos = int64(newNanoInt)
		}
		finalAmountType = datastruct.DEBIT

	} else if oldExpense.Type == datastruct.DEBIT && expenseData.Type == datastruct.CREDIT {
		// If the old expense was Debit but was updated to credit
		finalAmountUnits = int64(oldExpense.Amount.Units + expenseData.Amount.Units)
		finalAmountNanos = int64(oldExpense.Amount.Nanos + expenseData.Amount.Nanos)
		if finalAmountNanos > 999999999 {
			finalAmountUnits += 1
			nanoString := strconv.Itoa(int(finalAmountNanos))
			newNanoString := nanoString[1:]
			newNanoInt, _ := strconv.Atoi(newNanoString)
			finalAmountNanos = int64(newNanoInt)
		}
		finalAmountType = datastruct.CREDIT
	}
	// Update user balance
	err := es.dao.NewUserCollection().UpdateUserBalance(email, finalAmountUnits, int32(finalAmountNanos), finalAmountType)
	if err != nil {
		es.lg.Printf(err.Error())
		return nil, err
	}
	// Store expense in DB
	return es.dao.NewExpenseCollection().CreateExpense(expenseData)
}

func (es *expenseService) DeleteExpenseByID(userID, email, expenseID string) error {
	es.lg.Println("DeleteExpenseByID Called...")
	// Verification
	exp, err := es.GetExpenseByID(expenseID)
	if err != nil {
		return err
	}
	if exp.UserId != userID {
		return fmt.Errorf("not authorized")
	}
	var finalExpType datastruct.EXPENSE_TYPE
	if exp.Type == datastruct.CREDIT {
		finalExpType = datastruct.DEBIT
	} else if exp.Type == datastruct.DEBIT {
		finalExpType = datastruct.CREDIT
	}
	// Update the user balance
	err = es.dao.NewUserCollection().UpdateUserBalance(email, exp.Amount.Units, exp.Amount.Nanos, finalExpType)
	if err != nil {
		return err
	}
	return es.dao.NewExpenseCollection().DeleteExpenseByID(userID, expenseID)
}
