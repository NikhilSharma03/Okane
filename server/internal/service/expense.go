package service

import (
	"log"

	"github.com/NikhilSharma03/Okane/server/internal/repository"
)

// The ExpenseService interface defines methods to implement
type ExpenseService interface {
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
