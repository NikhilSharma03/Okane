package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/go-redis/redis/v8"
	"github.com/strongo/decimal"
)

// The UserCollection defines the methods a struct need to have
type UserCollection interface {
	CreateUser(id, name, email, password string) (*datastruct.User, error)
	GetUser(email string) (*datastruct.User, error)
	UserExists(email string) (bool, error)
	DeleteUser(email string) error
	UpdateUserBalance(email string, expenseAmountUnits int64, expenseAmountNanos int32, expenseType datastruct.EXPENSE_TYPE) error
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
	_, err = DB.HSet(context.Background(), USERS_COLLECTION, USERS+email, newUserJSON).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to store new user in redis hash")
	}

	// return created user
	return newUser, nil
}

func (*userCollection) GetUser(email string) (*datastruct.User, error) {
	// Fetch user from database
	data, err := DB.HGet(context.Background(), USERS_COLLECTION, USERS+email).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no user found with provided email")
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

func (*userCollection) UserExists(email string) (bool, error) {
	// Check if user exists in DB
	return DB.HExists(context.Background(), USERS_COLLECTION, USERS+email).Result()
}

func (*userCollection) DeleteUser(email string) error {
	// Remove user from DB
	_, err := DB.HDel(context.Background(), USERS_COLLECTION, USERS+email).Result()
	if err != nil {
		return fmt.Errorf("failed to remove user")
	}

	return nil
}

func (ur *userCollection) UpdateUserBalance(email string, expenseAmountUnits int64, expenseAmountNanos int32, expenseType datastruct.EXPENSE_TYPE) error {
	// Fetch the user with provided email ID
	user, err := ur.GetUser(email)
	if err != nil {
		return err
	}

	// Create float for calculation of Money type
	userBalaceFloat := decimal.NewDecimal64p2(user.Balance.Units, int8(user.Balance.Nanos))
	expenseAmountFloat := decimal.NewDecimal64p2(expenseAmountUnits, int8(expenseAmountNanos))
	// Calculate user result balance
	var resultBal decimal.Decimal64p2
	if expenseType == datastruct.CREDIT {
		resultBal = userBalaceFloat + expenseAmountFloat
	} else if expenseType == datastruct.DEBIT {
		resultBal = userBalaceFloat - expenseAmountFloat
	}
	resBal := resultBal.String()
	res := strings.Split(resBal, ".")
	userBalUnits, _ := strconv.Atoi(res[0])
	var userBalNanos int
	if len(res) == 2 {
		userBalNanos, _ = strconv.Atoi(res[1])
	} else {
		userBalNanos = 0
	}

	// Update user Balance in DB
	newUser := datastruct.NewUser(user.ID, user.Name, user.Email, user.Password)
	newUserBal := datastruct.NewBalance("USD", int64(userBalUnits), int32(userBalNanos))
	newUser.Balance = newUserBal
	// Marshalling newUser to JSON
	newUserJSON, err := newUser.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal new user")
	}
	// Storing in redis Hash
	_, err = DB.HSet(context.Background(), USERS_COLLECTION, USERS+user.Email, newUserJSON).Result()
	if err != nil {
		return fmt.Errorf("failed to store new user in redis hash")
	}
	return nil
}
