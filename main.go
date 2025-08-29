package main

import (
	"context"
	"log"
	"os"

	"net"

	"github.com/RyanHartadi06/clara-be/internal/handler"
	"github.com/RyanHartadi06/clara-be/pb/service"
	"github.com/RyanHartadi06/clara-be/pkg/database"
	"github.com/RyanHartadi06/clara-be/pkg/grpcmiddleware"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	godotenv.Load()
	ctx := context.Background()
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Panicf("Error when listening: %v", err)
	}
	defer lis.Close()

	database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Connected to database")

	serviceHandler := handler.NewServiceHandler()

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
		),
	)

	service.RegisterHelloWorldServiceServer(serv, serviceHandler)

	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Println("Reflection is Registered")
	}

	log.Println("Server is Running on port :50052")

	if err := serv.Serve(lis); err != nil {
		log.Panicf("Error when serving %v", err)
	}

}
