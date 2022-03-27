package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/youngvan/messenger-test/messenger/messengerpb"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	messengerServer := pb.Messenger.Server{}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
