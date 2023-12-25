package app

import (
	"context"
	"fmt"

	"github.com/NikhilSharma03/Okane/internal/datastruct"
	okanepb "github.com/NikhilSharma03/Okane/pkg/protobuf"
	"google.golang.org/genproto/googleapis/type/money"
)

// Update Expense updates the expense by the provided expense id
func (es *ExpenseService) UpdateExpenseByID(ctx context.Context, rq *okanepb.UpdateExpenseByIDRequest) (*okanepb.UpdateExpenseByIDResponse, error) {
	// Extract token from request
	token, ok := getCredFromMetadata(ctx, "token")
	if !ok {
		return nil, fmt.Errorf("token metadata not found in header")
	}
	// Extract userID and exp from jwt token
	tokenData, err := es.jwtService.ExtractDataFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	userID := fmt.Sprintf("%s", tokenData["id"])
	email := fmt.Sprintf("%s", tokenData["email"])
	if userID == "" || email == "" {
		return nil, fmt.Errorf("invalid token")
	}
	// Get expense ID
	expenseID := rq.GetId()
	if expenseID == "" {
		return nil, fmt.Errorf("no expense ID found")
	}
	// Fetch expense by ID
	expense, err := es.expenseService.GetExpenseByID(expenseID)
	if err != nil {
		return nil, err
	}
	// Verification
	if expense.UserId != userID {
		return nil, fmt.Errorf("not authorized")
	}
	// New Expense Data
	newExpenseID := expense.Id
	newExpenseUserID := expense.UserId
	newExpenseTitle := rq.GetExpenseData().GetTitle()
	newExpenseDescription := rq.GetExpenseData().GetDescription()
	newExpeseAmount := datastruct.NewAmount(expense.Amount.CurrencyCode, rq.GetExpenseData().GetAmount().GetUnits(), rq.GetExpenseData().GetAmount().GetNanos())
	newExpenseType := datastruct.EXPENSE_TYPE(rq.GetExpenseData().GetType())
	// Update expense
	newExpense, err := es.expenseService.UpdateExpenseByID(expense, newExpenseID, newExpenseUserID, email, newExpenseTitle, newExpenseDescription, newExpeseAmount, newExpenseType)
	if err != nil {
		return nil, err
	}
	return &okanepb.UpdateExpenseByIDResponse{
		Message: "Updated",
		ExpenseData: &okanepb.Expense{
			Id:          newExpense.Id,
			UserId:      newExpense.UserId,
			Title:       newExpense.Title,
			Description: newExpense.Description,
			Amount: &money.Money{
				CurrencyCode: newExpense.Amount.CurrencyCode,
				Units:        newExpense.Amount.Units,
				Nanos:        newExpense.Amount.Nanos,
			},
			Type: okanepb.Expense_EXPENSE_TYPE(newExpense.Type),
		},
	}, nil
}
