package app

import (
	"github.com/NikhilSharma03/Okane/server/internal/service"
	okanepb "github.com/NikhilSharma03/Okane/server/pkg/protobuf"
)

// The UserService implements handler of okanepb.UserService
// It is used in registering OkaneUserServer on grpc Server
type UserService struct {
	okanepb.UnimplementedOkaneUserServer
	userService service.UserService
	jwtService  service.JWTService
}

// NewUserService takes service.userService and returns UserService
func NewUserService(userService service.UserService, jwtService service.JWTService) *UserService {
	return &UserService{
		userService: userService,
		jwtService:  jwtService,
	}
}
