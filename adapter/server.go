package adapter

import (
	"context"

	"github.com/ksrnnb/chat-app-server/chatpb"
	"github.com/ksrnnb/chat-app-server/usecase"
)

type Server struct {
	chatpb.UnimplementedChatAppServiceServer
	UserInteractor usecase.IUserInteractor
}

func NewServer(i usecase.IUserInteractor) (*Server) {
	return &Server{UserInteractor: i}
}

func (s *Server) Login(ctx context.Context, req *chatpb.LoginRequest) (*chatpb.LoginResponse, error) {
	request := &usecase.LoginRequest{
		LoginId: req.GetLoginId(),
		Password: req.GetPassword(),
	}

	_, err := s.UserInteractor.GetUserByLoginId(request)

	if err != nil {
		return nil, err
	}

	newRes := &chatpb.LoginResponse{
		Message: "login succeeded",
	}

	return newRes, nil
}

func (s *Server) CreateMessage(ctx context.Context, req *chatpb.CreateMessageRequest) (*chatpb.CreateMessageResponse, error) {
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

func (s *Server) HelloMessage(ctx context.Context, req *chatpb.HelloRequest) (*chatpb.HelloResponse, error) {
	name := req.GetName()


	res := &chatpb.HelloResponse{
		Message: "Hello, " + name,
	}

	return res, nil
}