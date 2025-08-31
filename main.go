package main

import (
	"context"
	"log"
	"os"

	"net"

	"github.com/RyanHartadi06/clara-be/internal/grpcmiddleware"
	"github.com/RyanHartadi06/clara-be/internal/handler"
	"github.com/RyanHartadi06/clara-be/internal/repository"
	"github.com/RyanHartadi06/clara-be/internal/service"
	"github.com/RyanHartadi06/clara-be/pb/project"
	"github.com/RyanHartadi06/clara-be/pkg/database"
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

	db := database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Connected to database")

	projectRepository := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepository)
	projectHandler := handler.NewProjectHandler(projectService)

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
		),
	)

	project.RegisterProjectServiceServer(serv, projectHandler)

	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Println("Reflection is Registered")
	}

	log.Println("Server is Running on port :50052")

	if err := serv.Serve(lis); err != nil {
		log.Panicf("Error when serving %v", err)
	}

}
