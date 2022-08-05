package app

import (
	"context"
	"fmt"

	okanepb "github.com/NikhilSharma03/Okane/server/pkg/protobuf"
	"google.golang.org/genproto/googleapis/type/money"
)

// Get Expense By ID returns the expense by the provided expense id
func (es *ExpenseService) GetExpenseByID(ctx context.Context, rq *okanepb.GetExpenseByIDRequest) (*okanepb.GetExpenseByIDResponse, error) {
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
	if userID == "" {
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
	return &okanepb.GetExpenseByIDResponse{
		Message: "Fetched Expense",
		ExpensesData: &okanepb.Expense{
			Id:          expense.Id,
			UserId:      expense.UserId,
			Title:       expense.Title,
			Description: expense.Description,
			Amount: &money.Money{
				CurrencyCode: expense.Amount.CurrencyCode,
				Units:        expense.Amount.Units,
				Nanos:        expense.Amount.Nanos,
			},
			Type: okanepb.Expense_EXPENSE_TYPE(expense.Type),
		},
	}, nil
}
