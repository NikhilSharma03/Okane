package repository

import "github.com/NikhilSharma03/Okane/server/internal/datastruct"

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

	// Hashing the user password

	// Create new User
	newUser := datastruct.NewUser(,name, email, ,0)

	// Store user in hash database

	// Store userID in sorted set to get sorted data

	// return created user
	return newUser
}

func (*userCollection) GetUser() ([]*datastruct.User, error) {

}

func (*userCollection) GetUserByID(id, email string) (*datastruct.User, error) {

}

func (*userCollection) UpdateUserByID(id, email, name, password string) (*datastruct.User, error) {

}

func (*userCollection) DeleteUserByID(id, email string) (*datastruct.User, error) {

}
