package service

import (
	"log"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/NikhilSharma03/Okane/server/internal/repository"
)

// The UserService interface defines methods to implement
type UserService interface {
	CreateUser(name, email, password string) (*datastruct.User, error)
	GetUserByID(id, password string) (*datastruct.User, error)
	UpdateUserByID(id, password, name string) (*datastruct.User, error)
	DeleteUserByID(id, password string) (*datastruct.User, error)
}

// The userService struct take dao and logger (lg)
type userService struct {
	dao repository.DAO
	lg  *log.Logger
}

// The NewUserService takes dao and lg as parameter and returns UserService interface implemented struct
func NewUserService(dao repository.DAO, lg *log.Logger) UserService {
	return &userService{dao, lg}
}

func (us *userService) CreateUser(name, email, password string) (*datastruct.User, error)  {}
func (us *userService) GetUserByID(id, password string) (*datastruct.User, error)          {}
func (us *userService) UpdateUserByID(id, password, name string) (*datastruct.User, error) {}
func (us *userService) DeleteUserByID(id, password string) (*datastruct.User, error)       {}
