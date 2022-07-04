package repository

import "github.com/go-redis/redis/v8"

// The dao struct implement the DAO interface
type dao struct {
}

func NewDAO(db *redis.Client) *dao {
	DB = db
	return &dao{}
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
