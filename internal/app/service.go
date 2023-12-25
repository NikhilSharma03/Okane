package app

import (
	"github.com/NikhilSharma03/Okane/internal/service"
	okanepb "github.com/NikhilSharma03/Okane/pkg/protobuf"
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

// The ExpenseService implements handler of okanepb.ExpenseService
// It is used in registering OkaneExpenseServer on grpc Server
type ExpenseService struct {
	okanepb.UnimplementedOkaneExpenseServer
	expenseService service.ExpenseService
	jwtService     service.JWTService
}

// NewExpenseService takes service.expenseService and service.JWTService and returns ExpenseService
func NewExpenseService(expenseService service.ExpenseService, jwtService service.JWTService) *ExpenseService {
	return &ExpenseService{
		expenseService: expenseService,
		jwtService:     jwtService,
	}
}
