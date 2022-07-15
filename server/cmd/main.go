package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/NikhilSharma03/Okane/server/internal/app"
	"github.com/NikhilSharma03/Okane/server/internal/repository"
	"github.com/NikhilSharma03/Okane/server/internal/service"
	"github.com/NikhilSharma03/Okane/server/pkg/okanepb"
	"google.golang.org/grpc"
)

func main() {
	// Initialize logger
	lg := log.New(os.Stdout, "okane-api", log.LstdFlags)

	// Connect DB
	dbClient := repository.ConnectDB()
	// Check if DB is connected
	_, err := dbClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("Failed to connect DB.", err)
	}

	// Initialize DAO
	dao := repository.NewDAO(dbClient)
	// Register all services
	userService := service.NewUserService(dao, lg)

	// Starting gRPC server
	go func() {
		listener, err := net.Listen("tcp", "localhost:8001")
		if err != nil {
			log.Fatalln(err)
		}

		grpcServer := grpc.NewServer()
		okanepb.RegisterOkaneUserServer(grpcServer, app.NewUserService(userService))
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalln(err)
		}
	}()
}
