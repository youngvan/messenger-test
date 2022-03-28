package main

import (
	"context"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/jackc/pgx/v4/pgxpool"

	msgapp "messenger/pkg/messagesapp"

	pb "messenger/gen/proto/go/messagesproto/v1"
)

type MessengerServer struct {
	pb.UnimplementedMessengerServiceServer
	logger *zap.SugaredLogger
	db     *pgxpool.Pool
}

func (s *MessengerServer) Exchange(ctx context.Context, in *pb.ExchangeRequest) (*pb.ExchangeResponse, error) {
	s.logger.Info("Incoming request", zap.Any("request", in))

	// Validate AuthHash present
	if in.AuthHash == "" {
		return &pb.ExchangeResponse{
			Status:        pb.ExchangeResponse_STATUS_TYPE_AUTHFAIL,
			StatusMessage: "authHash missed"}, nil
	}

	// Create app
	app := msgapp.NewMessagesApp(s.logger, s.db)

	// Auth User
	// @todo implement native auth
	if err := app.Login(in.AuthHash); err != nil {
		return &pb.ExchangeResponse{
			Status:        pb.ExchangeResponse_STATUS_TYPE_AUTHFAIL,
			StatusMessage: "AUTH failed"}, nil
	}

	responce := &pb.ExchangeResponse{}

	// Save icoming messages to DB
	// @todo - how to handle error?
	app.SaveMessages(in.Messages)

	// Get new messages
	// @todo implement receiving messages from point of time
	// @todo - how to handle error?
	app.GetMessages(&responce.Messages)

	return responce, nil
}

func (s *MessengerServer) initLogger(logger *zap.SugaredLogger) error {
	s.logger = logger
	return nil
}

func (s *MessengerServer) initDb(db *pgxpool.Pool) error {
	s.db = db
	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	messengerServer := MessengerServer{}

	// Init logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugarZap := logger.Sugar()

	messengerServer.initLogger(sugarZap)

	// Init database
	// todo move connection string to ENV
	db, err := pgxpool.Connect(context.Background(), "postgres://messenger:00pass00@localhost:5432/messenger")
	if err != nil {
		log.Fatalf("failed start db: %v", err)
	}
	messengerServer.initDb(db)

	grpcServer := grpc.NewServer()

	pb.RegisterMessengerServiceServer(grpcServer, &messengerServer)

	if err := grpcServer.Serve(lis); err != nil {
		sugarZap.Error("failed to serve: %s", err)
	}
}