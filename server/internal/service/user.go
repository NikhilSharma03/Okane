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
	GetUser(email, password string) (*datastruct.User, error)
	UpdateUser(email, password, name string) (*datastruct.User, error)
	DeleteUser(email, password string) (*datastruct.User, error)
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
		us.lg.Printf("Failed to hash user password %+v", err.Error())
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

func (us *userService) GetUser(email, password string) (*datastruct.User, error) {
	// Log
	us.lg.Println("GetUser called...")

	// Fetch user by email
	user, err := us.dao.NewUserCollection().GetUser(email)
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

func (us *userService) UpdateUser(email, password, name string) (*datastruct.User, error) {
	// Log
	us.lg.Println("UpdateUser called...")

	// Fetch user by email (if exists)
	foundUser, err := us.dao.NewUserCollection().GetUser(email)
	if err != nil {
		us.lg.Printf("%+v", err.Error())
		return nil, fmt.Errorf("%+v", err.Error())
	}

	// Check if provided password is correct
	isPassCorrect := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if isPassCorrect != nil {
		us.lg.Printf("Incorrect password")
		return nil, fmt.Errorf("incorrect Password")
	}

	// Update user name
	foundUser.Name = name

	// Update User in DB
	user, err := us.dao.NewUserCollection().CreateUser(foundUser.ID, foundUser.Name, foundUser.Email, foundUser.Password)
	if err != nil {
		us.lg.Printf("Failed to update user")
		return nil, fmt.Errorf("failed to update user")
	}

	return user, nil
}

func (us *userService) DeleteUser(email, password string) (*datastruct.User, error) {
	us.lg.Println("DeleteUser called...")
	// Fetch user by email (if exists)
	foundUser, err := us.dao.NewUserCollection().GetUser(email)
	if err != nil {
		us.lg.Printf("%+v", err.Error())
		return nil, fmt.Errorf("%+v", err.Error())
	}

	// Check if provided password is correct
	isPassCorrect := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if isPassCorrect != nil {
		us.lg.Printf("Incorrect password")
		return nil, fmt.Errorf("incorrect Password")
	}

	// Delete User in DB
	err = us.dao.NewUserCollection().DeleteUser(foundUser.Email)
	if err != nil {
		us.lg.Printf("Failed to delete user")
		return nil, fmt.Errorf("failed to delete user")
	}

	return foundUser, nil
}
