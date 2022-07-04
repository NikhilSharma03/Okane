package repository

import "github.com/go-redis/redis/v8"

// DAO interface
type DAO interface {
	NewUserCollection() UserCollection
	NewExpenseCollection() ExpenseCollection
}

// The dao struct implement the DAO interface
type dao struct{}

// The NewDAO function returns a dao struct which implement DAO interface
func NewDAO(db *redis.Client) DAO {
	DB = db
	return &dao{}
}

func (*dao) NewUserCollection() UserCollection {
	return &userCollection{}
}

func (*dao) NewExpenseCollection() ExpenseCollection {
	return &expenseCollection{}
}

// The DB variable will contain the database client
// It is initialized when a new DAO is create
var DB *redis.Client

// ConnectDB connect to redis database
// It returns a new Redis Client
func ConnectDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
