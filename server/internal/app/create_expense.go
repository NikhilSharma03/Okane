package app

import (
	"context"
	"fmt"

	okanepb "github.com/NikhilSharma03/Okane/server/pkg/protobuf"
)

// Create Expense creates a new Expens
func (es *ExpenseService) CreateExpense(ctx context.Context, in *okanepb.CreateExpenseRequest) (*okanepb.CreateExpenseResponse, error) {
	// Get token from metadata
	token, ok := getCredFromMetadata(ctx, "token")
	if !ok {
		return nil, fmt.Errorf("token metadata not found in header")
	}
	// Extract userID, email and exp from jwt token
}
