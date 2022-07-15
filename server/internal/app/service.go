package app

import (
	"github.com/NikhilSharma03/Okane/server/internal/service"
	"github.com/NikhilSharma03/Okane/server/pkg/okanepb"
)

// The UserService implements handler of okanepb.UserService
// It is used in registering OkaneUserServer on grpc Server
type UserService struct {
	okanepb.UnimplementedOkaneUserServer
	userService service.UserService
}

// NewUserService takes service.userService and returns UserService
func NewUserService(userService service.UserService) *UserService {
	return &UserService{
		userService: userService,
	}
}
