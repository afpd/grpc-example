package main

import (
	"dev.home.arpa/devuser/grpc-example/rsocks"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting the server...")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := rsocks.Server{}
	grpcServer := grpc.NewServer()

	rsocks.RegisterTeleConnServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
