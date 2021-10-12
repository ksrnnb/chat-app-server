package app

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/ksrnnb/chat-app-server/adapter"
	"github.com/ksrnnb/chat-app-server/chatpb"
	"github.com/ksrnnb/chat-app-server/infrastructure/database"
	"github.com/ksrnnb/chat-app-server/infrastructure/repository"
	"github.com/ksrnnb/chat-app-server/usecase"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func Start() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sqlHandler := database.NewSqlHandler()
	repository := repository.NewUserRepository(sqlHandler)
	interactor := usecase.NewUserInteractor(repository)
	server := adapter.NewServer(interactor)

	chatpb.RegisterChatAppServiceServer(s, server)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}