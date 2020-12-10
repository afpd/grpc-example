package main

import (
	"fmt"

	"dev.home.arpa/devuser/grpc-example/rsocks"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Starting the client")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := rsocks.NewTeleConnClient(conn)
	stream, err := c.TeleConn(context.Background())

	go func() {
		for {
			br, _ := stream.Recv()
			log.Printf("Received: %v", string(br.Body))
		}
	}()
	go func() {
		for {
			stream.Send(&rsocks.Message{Body: []byte("xlam")})
		}
	}()
	done := make(chan bool)
	<-done

}
