package app

import (
	"context"
	"fmt"

	okanepb "github.com/NikhilSharma03/Okane/server/pkg/protobuf"
	"google.golang.org/genproto/googleapis/type/money"
)

// Get Expense returns all the expenses by the provided user id
func (es *ExpenseService) GetExpense(ctx context.Context, rq *okanepb.GetExpenseRequest) (*okanepb.GetExpenseResponse, error) {
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
	// Verification
	if userID != rq.GetUserId() {
		return nil, fmt.Errorf("userID did not match. invalid userID")
	}

	// Fetch all expenses with UserID
	expenses, err := es.expenseService.GetExpenses(userID)
	if err != nil {
		return nil, err
	}
	// Generate response data
	var responseExpenseData []*okanepb.Expense
	for _, item := range expenses {
		exp := &okanepb.Expense{
			Id:          item.Id,
			UserId:      item.UserId,
			Title:       item.Title,
			Description: item.Description,
			Amount: &money.Money{
				CurrencyCode: item.Amount.CurrencyCode,
				Units:        item.Amount.Units,
				Nanos:        item.Amount.Nanos,
			},
			Type: okanepb.Expense_EXPENSE_TYPE(item.Type),
		}
		responseExpenseData = append(responseExpenseData, exp)
	}
	return &okanepb.GetExpenseResponse{
		Message:      "Fetched expenses Successfully",
		ExpensesData: responseExpenseData,
	}, nil
}
