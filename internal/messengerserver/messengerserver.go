package messengerserver

import (
	"context"
	"messenger/pkg/messagesapp"
	"messenger/pkg/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "messenger/proto/gen/proto/go/messagesproto/v1"
)

type MessengerServer struct {
	pb.UnimplementedMessengerServiceServer
	App         *messagesapp.MessagesApp
	CurrentUser *user.UserStruct
}

func (s *MessengerServer) Exchange(ctx context.Context, in *pb.ExchangeRequest) (*pb.ExchangeResponse, error) {

	// Validate AuthHash present
	if in.AuthHash == "" {
		return nil, status.Errorf(codes.Unauthenticated, "authHash missed")
	}

	// Auth User
	// @todo implement native auth
	if err := s.CurrentUser.Login(ctx, in.AuthHash); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Auth missed")
	}

	responce := pb.ExchangeResponse{}

	// Save icoming messages to DB
	if err := s.App.SaveMessages(ctx, s.CurrentUser, in.Messages); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Error with saving messages")
	}

	// Get new messages
	// @todo implement receiving messages from point of time
	messages, err := s.App.GetMessages(ctx, s.CurrentUser)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Error with getting new")
	}

	for _, msg := range messages {
		responce.Messages = append(responce.Messages, &pb.Message{
			Subject: msg.Subject,
			Body:    msg.Body})
	}

	return &responce, nil
}
