package app

import (
	"context"
	"fmt"

	okanepb "github.com/NikhilSharma03/Okane/server/pkg/protobuf"
	"google.golang.org/genproto/googleapis/type/money"
)

// CreateUser creates a new user
func (us *UserService) CreateUser(ctx context.Context, req *okanepb.CreateUserRequest) (*okanepb.CreateUserResponse, error) {
	// Get User Data from Request
	uData := req.GetUserData()
	// Create User using User Service
	user, err := us.userService.CreateUser(uData.Name, uData.Email, uData.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create user %v", err.Error())
	}
	// If user created successfully
	userResponse := &okanepb.CreateUserResponse{Message: "User Created Successfully", UserData: &okanepb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Balance: &money.Money{
			CurrencyCode: user.Balance.CurrencyCode,
			Units:        user.Balance.Units,
			Nanos:        user.Balance.Nanos,
		},
	}}
	return userResponse, nil
}
