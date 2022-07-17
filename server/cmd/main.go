package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/NikhilSharma03/Okane/server/internal/app"
	"github.com/NikhilSharma03/Okane/server/internal/repository"
	"github.com/NikhilSharma03/Okane/server/internal/service"
	okanepb "github.com/NikhilSharma03/Okane/server/pkg/protobuf"
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
		// Initialize Listener at Port 8001
		listener, err := net.Listen("tcp", "localhost:8001")
		if err != nil {
			log.Fatalln(err)
		}
		// Initialize new grpc server
		grpcServer := grpc.NewServer()
		// Register server
		okanepb.RegisterOkaneUserServer(grpcServer, app.NewUserService(userService))
		// Server grpc server on listener
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalln(err)
		}
	}()
}
