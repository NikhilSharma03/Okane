package service

import (
	"fmt"
	"log"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/NikhilSharma03/Okane/server/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (us *userService) CreateUser(name, email, password string) (*datastruct.User, error) {
	// Log
	us.lg.Println("CreateUser called...")

	// Generating new ID for user
	id := uuid.New().String()
	if id == "" {
		us.lg.Printf("Failed to generate id")
		return nil, fmt.Errorf("failed to generate user ID")
	}

	// Hashing the user password
	hpassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		us.lg.Printf("Failed to hash user password %v", err.Error())
		return nil, fmt.Errorf("failed to hash user password")
	}

	// create user
	user, err := us.dao.NewUserCollection().CreateUser(id, name, email, string(hpassword))
	if err != nil {
		us.lg.Printf("Failed to create user")
		return nil, fmt.Errorf("failed to create user")
	}

	return user, nil
}

func (us *userService) GetUserByID(id, password string) (*datastruct.User, error) {
	// Log
	us.lg.Println("GetUserByID called...")

	// Fetch user by id
	user, err := us.dao.NewUserCollection().GetUserByID(id, password)
	if err != nil {
		us.lg.Printf("%+v", err.Error())
		return nil, fmt.Errorf("%+v", err.Error())
	}

	// Check if provided password is correct
	isPassCorrect := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if isPassCorrect != nil {
		us.lg.Printf("Incorrect password")
		return nil, fmt.Errorf("incorrect Password")
	}

	return user, nil
}
func (us *userService) UpdateUserByID(id, password, name string) (*datastruct.User, error) {}
func (us *userService) DeleteUserByID(id, password string) (*datastruct.User, error)       {}
