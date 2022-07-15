package main

import (
	"context"
	"log"
	"os"

	"github.com/NikhilSharma03/Okane/server/internal/repository"
)

func main() {
	// Initialize logger
	lg := log.New(os.Stdout, "okane-api", log.LstdFlags)

	// Connect DB
	dbClient := repository.ConnectDB()
	// Check if DB is connected
	_, err := dbClient.Ping(context.Background()).Result()
	if err != nil {
		lg.Printf("Failed to connect DB. %+v", err.Error())
	}

	// Initialize DAO
	dao := repository.NewDAO(dbClient)

}
