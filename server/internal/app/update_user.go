package app

import (
	"context"
	"fmt"

	okanepb "github.com/NikhilSharma03/Okane/server/pkg/protobuf"
	"google.golang.org/genproto/googleapis/type/money"
)

// UpdateUserByID updates user (if exists) with provided ID
func (us *UserService) UpdateUserByID(ctx context.Context, req *okanepb.UpdateUserByIDRequest) (*okanepb.UpdateUserByIDResponse, error) {
	// Get id, name(to update), password from request
	id, name, pass := req.GetId(), req.GetName(), req.GetPassword()
	// Update user using service
	updatedUser, err := us.userService.UpdateUserByID(id, pass, name)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &okanepb.UpdateUserByIDResponse{
		Message: "Updated User Successfully",
		UserData: &okanepb.User{
			Id:       updatedUser.ID,
			Name:     updatedUser.Name,
			Email:    updatedUser.Email,
			Password: updatedUser.Password,
			Balance: &money.Money{
				CurrencyCode: updatedUser.Balance.CurrencyCode,
				Units:        updatedUser.Balance.Units,
				Nanos:        updatedUser.Balance.Nanos,
			},
		},
	}, nil
}
