package app

import (
	"context"
	"fmt"

	okanepb "github.com/NikhilSharma03/Okane/server/pkg/protobuf"
	"google.golang.org/genproto/googleapis/type/money"
)

// GetUserByID fetch and returns user if exists
func (us *UserService) GetUserByID(ctx context.Context, req *okanepb.GetUserByIDRequest) (*okanepb.GetUserByIDResponse, error) {
	// Get UserID from request
	userID := req.GetId()
	// Get Password from metadata
	userCred, ok := getCredFromMetadata(ctx)
	if !ok {
		return nil, fmt.Errorf("cred metadata not found in header")
	}
	// Check and return user if exists
	userData, err := us.userService.GetUserByID(userID, userCred)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &okanepb.GetUserByIDResponse{
		Message: "Found User",
		UserData: &okanepb.User{
			Id:       userData.ID,
			Name:     userData.Name,
			Email:    userData.Email,
			Password: userData.Password,
			Balance: &money.Money{
				CurrencyCode: userData.Balance.CurrencyCode,
				Units:        userData.Balance.Units,
				Nanos:        userData.Balance.Nanos,
			},
		},
	}, nil
}
