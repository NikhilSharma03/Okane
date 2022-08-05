package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/NikhilSharma03/Okane/server/internal/datastruct"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
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
	MAX_NANOS        = 999999999
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
		return nil, fmt.Errorf("failed to unmarshal found user %v", err)
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
	// Calculate user result balance
	var resultBalanceUnits int64
	var resultBalanceNanos int64
	if expenseType == datastruct.CREDIT {
		userBalUnits := strconv.Itoa(int(user.Balance.Units))
		userBalNanos := strconv.Itoa(int(user.Balance.Nanos))
		expenseAmtUnits := strconv.Itoa(int(expenseAmountUnits))
		expenseAmtNanos := strconv.Itoa(int(expenseAmountNanos))
		userBal := userBalUnits + "."
		for i := len(userBalNanos); i < 9; i++ {
			userBal += "0"
		}
		if strings.Contains(userBalNanos, "-") {
			userBal += strings.Replace(userBalNanos, "-", "", 1)
			if !strings.Contains(userBal, "-") {
				temp := "-" + userBal
				userBal = temp
			}
		} else {
			userBal += userBalNanos
		}
		expAmt := expenseAmtUnits + "."
		for i := len(expenseAmtNanos); i < 9; i++ {
			expAmt += "0"
		}
		if strings.Contains(expenseAmtNanos, "-") {
			expAmt += strings.Replace(expenseAmtNanos, "-", "", 1)
			if !strings.Contains(expAmt, "-") {
				temp := "-" + expAmt
				expAmt = temp
			}
		} else {
			expAmt += expenseAmtNanos
		}
		decimal.DivisionPrecision = 9
		userBalDec, err := decimal.NewFromString(userBal)
		if err != nil {
			return err
		}
		expAmtDec, err := decimal.NewFromString(expAmt)
		if err != nil {
			return err
		}
		result := userBalDec.Add(expAmtDec).String()
		fmt.Println("Result ====")
		fmt.Println(result)
		resultArr := strings.Split(result, ".")
		// Set Nanos and Units
		resultUnit, _ := strconv.Atoi(string(resultArr[0]))
		resultBalanceUnits = int64(resultUnit)
		if len(resultArr) == 2 {
			currNanos := string(resultArr[1])
			currLength := len(currNanos)
			for currLength < 9 {
				currNanos += "0"
				currLength++
			}
			resultNanos, _ := strconv.Atoi(currNanos)
			resultBalanceNanos = int64(resultNanos)
			if strings.Contains(result, "-") {
				temp := -(resultBalanceNanos)
				resultBalanceNanos = temp
			}
		} else {
			resultBalanceNanos = 0
		}
	} else if expenseType == datastruct.DEBIT {
		userBalUnits := strconv.Itoa(int(user.Balance.Units))
		userBalNanos := strconv.Itoa(int(user.Balance.Nanos))
		expenseAmtUnits := strconv.Itoa(int(expenseAmountUnits))
		expenseAmtNanos := strconv.Itoa(int(expenseAmountNanos))
		userBal := userBalUnits + "."
		for i := len(userBalNanos); i < 9; i++ {
			userBal += "0"
		}
		if strings.Contains(userBalNanos, "-") {
			userBal += strings.Replace(userBalNanos, "-", "", 1)
			if !strings.Contains(userBal, "-") {
				temp := "-" + userBal
				userBal = temp
			}
		} else {
			userBal += userBalNanos
		}
		expAmt := expenseAmtUnits + "."
		for i := len(expenseAmtNanos); i < 9; i++ {
			expAmt += "0"
		}
		if strings.Contains(expenseAmtNanos, "-") {
			expAmt += strings.Replace(expenseAmtNanos, "-", "", 1)
			if !strings.Contains(expAmt, "-") {
				temp := "-" + expAmt
				expAmt = temp
			}
		} else {
			expAmt += expenseAmtNanos
		}
		decimal.DivisionPrecision = 9
		userBalDec, err := decimal.NewFromString(userBal)
		if err != nil {
			return err
		}
		expAmtDec, err := decimal.NewFromString(expAmt)
		if err != nil {
			return err
		}
		result := userBalDec.Sub(expAmtDec).String()
		resultArr := strings.Split(result, ".")
		// Set Nanos and Units
		resultUnit, _ := strconv.Atoi(string(resultArr[0]))
		resultBalanceUnits = int64(resultUnit)
		if len(resultArr) == 2 {
			currNanos := string(resultArr[1])
			currLength := len(currNanos)
			for currLength < 9 {
				currNanos += "0"
				currLength++
			}
			resultNanos, _ := strconv.Atoi(currNanos)
			resultBalanceNanos = int64(resultNanos)
			if strings.Contains(result, "-") {
				temp := -(resultBalanceNanos)
				resultBalanceNanos = temp
			}
		} else {
			resultBalanceNanos = 0
		}
	}
	// // Update user Balance in DB
	newUser := datastruct.NewUser(user.ID, user.Name, user.Email, user.Password)
	newUserBal := datastruct.NewBalance("USD", int64(resultBalanceUnits), int32(resultBalanceNanos))
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
