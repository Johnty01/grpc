package main

import (
	chat "grpc/chat/Users/john/grpc/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":7890")
	if err != nil {
		log.Fatalf("Failed to listen on port 7890: %v", err)
	}

	server := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server grpc server over port 7890: %v", err)
	}
}
