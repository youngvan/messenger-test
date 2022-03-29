package messengerserver

import (
	"context"
	"messenger/pkg/messagesapp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "messenger/proto/gen/proto/go/messagesproto/v1"
)

type MessengerServer struct {
	pb.UnimplementedMessengerServiceServer
	App *messagesapp.MessagesApp
}

func (s *MessengerServer) Exchange(ctx context.Context, in *pb.ExchangeRequest) (*pb.ExchangeResponse, error) {

	var requestMessages []*messagesapp.MessagesStruct
	for _, msg := range in.Messages {
		requestMessages = append(requestMessages, &messagesapp.MessagesStruct{
			Subject:      msg.Subject,
			Body:         msg.Body,
			ReceiverHash: msg.ReceiverHash})
	}

	// Save icoming messages to DB
	if err := s.App.SaveMessages(ctx, requestMessages); err != nil {
		return nil, status.Errorf(codes.Internal, "Error with saving messages")
	}

	// Get new messages
	// @todo implement receiving messages from point of time
	messages, err := s.App.GetMessages(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error with getting new messages")
	}

	var responce pb.ExchangeResponse
	for _, msg := range messages {
		responce.Messages = append(responce.Messages, &pb.Message{
			Subject: msg.Subject,
			Body:    msg.Body})
	}

	return &responce, nil
}
