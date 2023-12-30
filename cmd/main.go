package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/NikhilSharma03/Okane/internal/app"
	"github.com/NikhilSharma03/Okane/internal/repository"
	"github.com/NikhilSharma03/Okane/internal/service"
	okanepb "github.com/NikhilSharma03/Okane/pkg/protobuf"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Initialize environmental variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize logger
	lg := log.New(os.Stdout, "okane-api", log.LstdFlags)

	// Connect DB
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPass := os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalln("Failed to get REDIS_DB")
	}
	dbClient := repository.ConnectDB(redisAddress, redisPass, redisDB)
	// Check if DB is connected
	_, err = dbClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("Failed to connect DB.", err)
	}

	// Initialize DAO
	dao := repository.NewDAO(dbClient)
	// Register all services
	userService := service.NewUserService(dao, lg)
	expenseService := service.NewExpenseService(dao, lg)
	jwtService := service.NewJWTService(lg)

	// Starting gRPC server
	go func() {
		// Initialize Listener at Port 8001
		listener, err := net.Listen("tcp", ":8001")
		if err != nil {
			log.Fatalln(err)
		}
		// Initialize new grpc server
		grpcServer := grpc.NewServer()
		// Register server
		okanepb.RegisterOkaneUserServer(grpcServer, app.NewUserService(userService, jwtService))
		okanepb.RegisterOkaneExpenseServer(grpcServer, app.NewExpenseService(expenseService, jwtService))
		// Server grpc server on listener
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Client connection to the gRPC server
	conn, err := grpc.DialContext(
		context.Background(),
		":8001",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register OkaneUser
	err = okanepb.RegisterOkaneUserHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register user handler:", err)
	}
	err = okanepb.RegisterOkaneExpenseHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register expense handler:", err)
	}

	port := ":8000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}

	gwServer := &http.Server{
		Addr:    port,
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on", port)
	log.Fatalln(gwServer.ListenAndServe())
}
