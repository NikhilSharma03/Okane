package app

import (
	"context"
	"fmt"

	okanepb "github.com/NikhilSharma03/Okane/server/pkg/protobuf"
	"google.golang.org/genproto/googleapis/type/money"
)

// DeleteUserByID deletes the user with provided ID
func (us *UserService) DeleteUserByID(ctx context.Context, req *okanepb.DeleteUserByIDRequest) (*okanepb.DeleteUserByIDResponse, error) {
	// Get Email from request
	userEmail := req.GetEmail()
	// Get Password from metadata
	userCred, ok := getCredFromMetadata(ctx, "cred")
	if !ok {
		return nil, fmt.Errorf("cred metadata not found in header")
	}
	// Check and delete user if exists
	userData, err := us.userService.DeleteUser(userEmail, userCred)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &okanepb.DeleteUserByIDResponse{
		Message: "Deleted User",
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
