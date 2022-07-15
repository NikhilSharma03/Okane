package app

import (
	"github.com/NikhilSharma03/Okane/server/internal/service"
	"github.com/NikhilSharma03/Okane/server/pkg/okanepb"
)

type UserService struct {
	okanepb.UnimplementedOkaneUserServer
	userService service.UserService
}

func NewUserService(userService service.UserService) *UserService {
	return &UserService{
		userService: userService,
	}
}
