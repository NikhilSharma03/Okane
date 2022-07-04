package repository

import "github.com/NikhilSharma03/Okane/server/internal/datastruct"

// The UserCollection defines the methods a struct need to have
type UserCollection interface {
	CreateUser()
	GetUser() []*datastruct.User
	GetUserByID(id, email string) *datastruct.User
	UpdateUserByID()
	DeleteUserByID()
}

// The userCollection struct implements method as defined in UserCollection interface
type userCollection struct{}
