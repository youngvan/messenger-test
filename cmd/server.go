package main

import (
	"context"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/jackc/pgx/v4/pgxpool"

	"messenger/pkg/messagesapp"

	"messenger/internal/messengerserver"
	"messenger/pkg/user"

	pb "messenger/proto/gen/proto/go/messagesproto/v1"
)

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

	// question: who should crteate UserApp?
	// question: server.go and pass to rpcServer? rpcServer? or here?

	// question: chains of passing db?

	// question - too many possible objects in main?

	currentUser := user.UserStruct{
		Db:     app.Db,
		Logger: app.Logger,
	}

	// question: duplication of prefit and name?
	// question: module folder, module file name (lowercase?), and module name corellation
	messengerServer := messengerserver.MessengerServer{
		App:         &app,
		CurrentUser: &currentUser,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMessengerServiceServer(grpcServer, &messengerServer)

	if err := grpcServer.Serve(lis); err != nil {
		sugarZap.Error("failed to serve: %s", err)
	}
}
