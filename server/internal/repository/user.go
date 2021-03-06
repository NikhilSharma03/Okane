package repository

import (
	"context"
	"fmt"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/go-redis/redis/v8"
)

// The UserCollection defines the methods a struct need to have
type UserCollection interface {
	CreateUser(id, name, email, password string) (*datastruct.User, error)
	GetUserByID(id string) (*datastruct.User, error)
	DeleteUserByID(id string) error
}

// User Database constants
const (
	USERS_COLLECTION = "users"
	USERS            = "users:"
)

// The userCollection struct implements method as defined in UserCollection interface
type userCollection struct{}

func (*userCollection) CreateUser(id, name, email, password string) (*datastruct.User, error) {
	// Create new User
	newUser := datastruct.NewUser(id, name, email, password)
	newUserBal := datastruct.NewBalance("USD", 0, 0)
	newUser.Balance = newUserBal

	// Store user in hash database
	// Marshalling newUser to JSON
	newUserJSON, err := newUser.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal new user")
	}

	// Storing in redis Hash
	_, err = DB.HSet(context.Background(), USERS_COLLECTION, USERS+id, newUserJSON).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to store new user in redis hash")
	}

	// return created user
	return newUser, nil
}

func (*userCollection) GetUserByID(id string) (*datastruct.User, error) {
	// Fetch user from database
	data, err := DB.HGet(context.Background(), USERS_COLLECTION, USERS+id).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no user found with provided ID")
		}
		return nil, fmt.Errorf("failed to fetch user")
	}

	// Unmarshal the found string data to User struct
	var foundUser datastruct.User
	err = foundUser.Unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal found user")
	}

	// If all correct return the found user
	return &foundUser, nil
}

func (*userCollection) DeleteUserByID(id string) error {
	// Remove user from DB
	_, err := DB.HDel(context.Background(), USERS_COLLECTION, USERS+id).Result()
	if err != nil {
		return fmt.Errorf("failed to remove user")
	}

	return nil
}
