package repository

import (
	"context"
	"fmt"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// The UserCollection defines the methods a struct need to have
type UserCollection interface {
	CreateUser(name, email, password string) (*datastruct.User, error)
	GetUserByEmail(id, password string) (*datastruct.User, error)
	UpdateUserByEmail(id, password, name string) (*datastruct.User, error)
	DeleteUserByEmail(id, password string) (*datastruct.User, error)
}

// User Database constants
const (
	USERS_COLLECTION = "users"
	USERS            = "users:"
)

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
	_, err = DB.HSet(context.Background(), USERS_COLLECTION, USERS+id, newUserJSON).Result()
	if err != nil {
		return nil, fmt.Errorf("Failed to store new user in redis hash")
	}

	// return created user
	return newUser, nil
}

func (*userCollection) GetUserByID(id, password string) (*datastruct.User, error) {
	// Fetch user from database
	data, err := DB.HGet(context.Background(), USERS_COLLECTION, USERS+id).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("No user found with provided ID")
		}
		return nil, fmt.Errorf("Failed to fetch user")
	}

	// Unmarshal the found string data to User struct
	var foundUser datastruct.User
	err = foundUser.Unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal found user")
	}

	// Check if provided password is correct
	isPassCorrect := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if isPassCorrect != nil {
		return nil, fmt.Errorf("Incorrect Password")
	}

	// If all correct return the found user
	return &foundUser, nil
}

func (*userCollection) UpdateUserByID(id, password, name string) (*datastruct.User, error) {
	// Fetch user from database
	data, err := DB.HGet(context.Background(), USERS_COLLECTION, USERS+id).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("No user found with provided ID")
		}
		return nil, fmt.Errorf("Failed to fetch user")
	}

	// Unmarshal the found string data to User struct
	var foundUser datastruct.User
	err = foundUser.Unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal found user")
	}

	// Check if provided password is correct
	isPassCorrect := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if isPassCorrect != nil {
		return nil, fmt.Errorf("Incorrect Password")
	}

	// Update user name
	foundUser.Name = name

	// Update user in the DB
	// Marshalling found User to JSON
	updatedUser, err := foundUser.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal new user")
	}

	// Storing in redis Hash
	_, err = DB.HSet(context.Background(), USERS_COLLECTION, USERS+id, updatedUser).Result()
	if err != nil {
		return nil, fmt.Errorf("Failed to store updated user in redis hash")
	}

	// If all correct return the found user
	return &foundUser, nil
}

func (*userCollection) DeleteUserByID(id, password string) (*datastruct.User, error) {

}
