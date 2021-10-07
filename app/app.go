package app

import (
	"context"
	"log"
	"net"

	"github.com/ksrnnb/chat-app-server/chatpb"
	"google.golang.org/grpc"
)

const (
	// port = ":50051"
	port = ":9000"
)

type server struct {
	chatpb.UnimplementedChatAppServiceServer
}

type Message struct {
	Message string
	SenderId string
	ReceiverId string
	
}

func Start() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	chatpb.RegisterChatAppServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateMessage(ctx context.Context, req *chatpb.CreateMessageRequest) (*chatpb.CreateMessageResponse, error) {
	message := req.GetMessage()

	// TODO: save message

	messageItem := &chatpb.Message{
		Id: "123abc",
		SenderId: message.SenderId,
		ReceiverId: message.ReceiverId,
		SentAt: message.SentAt,
	}

	res := &chatpb.CreateMessageResponse{
		Message: messageItem,
	}

	return res, nil
}

func (s *server) HelloMessage(ctx context.Context, req *chatpb.HelloRequest) (*chatpb.HelloResponse, error) {
	name := req.GetName()


	res := &chatpb.HelloResponse{
		Message: "Hello, " + name,
	}

	return res, nil
}