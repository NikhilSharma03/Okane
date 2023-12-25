package app

import (
	"context"
	"fmt"

	okanepb "github.com/NikhilSharma03/Okane/pkg/protobuf"
	"google.golang.org/genproto/googleapis/type/money"
)

// GetUserByID fetch and returns user if exists
func (us *UserService) GetUserByID(ctx context.Context, req *okanepb.GetUserByIDRequest) (*okanepb.GetUserByIDResponse, error) {
	// Get UserEmail from request
	userEmail := req.GetEmail()
	// Get Password from metadata
	userCred, ok := getCredFromMetadata(ctx, "cred")
	if !ok {
		return nil, fmt.Errorf("cred metadata not found in header")
	}
	// Check and get user if exists
	userData, err := us.userService.GetUser(userEmail, userCred)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	// Generate JWT token
	token, err := us.jwtService.GenerateJWT(userData.ID, userData.Email)
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
		Token: token,
	}, nil
}
