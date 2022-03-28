package messengerserver

type MessengerServer struct {
	pb.UnimplementedMessengerServiceServer
	App *messagesapp.MessagesApp
}

func (s *MessengerServer) Exchange(ctx context.Context, in *pb.ExchangeRequest) (*pb.ExchangeResponse, error) {

	// Validate AuthHash present
	if in.AuthHash == "" {
		return nil, status.Errorf(codes.Unauthenticated, "authHash missed")
	}

	// Auth User
	// @todo implement native auth
	if err := s.App.Login(in.AuthHash); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Auth missed")
	}

	responce := &pb.ExchangeResponse{}

	// Save icoming messages to DB
	if err := s.App.SaveMessages(in.Messages); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Error with saving messages")
	}

	// Get new messages
	// @todo implement receiving messages from point of time
	if responce.Messages, err := s.App.GetMessages(); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Error with getting new")
	}

	return responce, nil
}