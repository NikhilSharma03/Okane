package app

import (
	"context"
	"fmt"

	"github.com/NikhilSharma03/Okane/internal/datastruct"
	okanepb "github.com/NikhilSharma03/Okane/pkg/protobuf"
	"google.golang.org/genproto/googleapis/type/money"
)

// Create Expense creates a new Expens
func (es *ExpenseService) CreateExpense(ctx context.Context, rq *okanepb.CreateExpenseRequest) (*okanepb.CreateExpenseResponse, error) {
	// Get token from metadata
	token, ok := getCredFromMetadata(ctx, "token")
	if !ok {
		return nil, fmt.Errorf("token metadata not found in header")
	}
	// Extract userID, email and exp from jwt token
	tokenData, err := es.jwtService.ExtractDataFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	userEmail := fmt.Sprintf("%s", tokenData["email"])
	userID := fmt.Sprintf("%s", tokenData["id"])
	if userEmail == "" || userID == "" {
		return nil, fmt.Errorf("failed to extract email and id from token")
	}
	// User request data
	title := rq.GetExpenseData().GetTitle()
	description := rq.GetExpenseData().GetDescription()
	amount := datastruct.NewAmount("USD", rq.GetExpenseData().GetAmount().GetUnits(), rq.ExpenseData.GetAmount().GetNanos())
	expenseType := rq.GetExpenseData().GetType()

	// Create a new Expense
	expense, err := es.expenseService.CreateExpense(userID, userEmail, title, description, amount, datastruct.EXPENSE_TYPE(expenseType))
	if err != nil {
		return nil, err
	}
	return &okanepb.CreateExpenseResponse{
		Message: "Expense create successfully",
		ExpenseData: &okanepb.Expense{
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
