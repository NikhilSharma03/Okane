package repository

import (
	"context"
	"fmt"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// The UserCollection defines the methods a struct need to have
type UserCollection interface {
	CreateUser(name, email, password string) (*datastruct.User, error)
	GetUser() ([]*datastruct.User, error)
	GetUserByID(id, email string) (*datastruct.User, error)
	UpdateUserByID(id, email, name, password string) (*datastruct.User, error)
	DeleteUserByID(id, email string) (*datastruct.User, error)
}

// User Database constants
const USERS = "users:"

// The userCollection struct implements method as defined in UserCollection interface
type userCollection struct{}

func (*userCollection) CreateUser(name, email, password string) (*datastruct.User, error) {
	// Generating new ID for user
	id := uuid.New().String()
	if id == "" {
		return nil, fmt.Errorf("Failed to generate user ID")
	}

	// Hashing the user password
	hpassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, fmt.Errorf("Failed to hash user password")
	}

	// Create new User
	newUser := datastruct.NewUser(id, name, email, string(hpassword), 0)

	// Store user in hash database
	// Marshalling newUser to JSON
	newUserJSON, err := newUser.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal new user")
	}

	// Storing in redis Hash
	_, err = DB.HSet(context.Background(), USERS+id, newUserJSON).Result()
	if err != nil {
		return nil, fmt.Errorf("Failed to store new user in redis hash")
	}

	// return created user
	return newUser, nil
}

func (*userCollection) GetUser() ([]*datastruct.User, error) {

}

func (*userCollection) GetUserByID(id, email string) (*datastruct.User, error) {

}

func (*userCollection) UpdateUserByID(id, email, name, password string) (*datastruct.User, error) {

}

func (*userCollection) DeleteUserByID(id, email string) (*datastruct.User, error) {

}
