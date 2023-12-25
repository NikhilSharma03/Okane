package app

import (
	"context"
	"fmt"

	okanepb "github.com/NikhilSharma03/Okane/pkg/protobuf"
)

// Delete Expense deletes the expense by the provided expense id
func (es *ExpenseService) DeleteExpenseByID(ctx context.Context, rq *okanepb.DeleteExpenseByIDRequest) (*okanepb.DeleteExpenseByIDResponse, error) {
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
	// Get expense ID
	expenseID := rq.GetId()
	if expenseID == "" {
		return nil, fmt.Errorf("no expense ID found")
	}
	// Delete expense using service
	err = es.expenseService.DeleteExpenseByID(userID, userEmail, expenseID)
	if err != nil {
		return nil, err
	}
	return &okanepb.DeleteExpenseByIDResponse{
		Message: "Expense deleted successfully",
	}, nil
}
