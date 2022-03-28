package main

import (
	"context"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jackc/pgx/v4/pgxpool"

	"messenger/pkg/messagesapp"

	pb "messenger/gen/proto/go/messagesproto/v1"
)

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

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Init logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugarZap := logger.Sugar()

	// Init database
	// todo move connection string to ENV
	dbConn, err := pgxpool.Connect(context.Background(), "postgres://messenger:00pass00@localhost:5432/messenger")
	if err != nil {
		log.Fatalf("failed start db: %v", err)
	}

	// Create app
	app := messagesapp.MessagesApp{
		Logger: sugarZap,
		Db:     dbConn,
	}

	messengerServer := MessengerServer{
		App: &app,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMessengerServiceServer(grpcServer, &messengerServer)

	if err := grpcServer.Serve(lis); err != nil {
		sugarZap.Error("failed to serve: %s", err)
	}
}
